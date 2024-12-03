package sure_pkg_gen

import (
	"go/ast"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/yyle88/done"
	"github.com/yyle88/formatgo"
	"github.com/yyle88/rese"
	"github.com/yyle88/sure"
	"github.com/yyle88/sure/internal/utils"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_aktnorm"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_astnode"
	"github.com/yyle88/syntaxgo/syntaxgo_search"
	"github.com/yyle88/tern"
	"github.com/yyle88/tern/zerotern"
)

func GenerateSurePackage(
	t *testing.T,
	sourceRoot string,
	errorHandlingMode sure.ErrorHandlingMode,
	sourcePackagePath string,
) {
	GenerateSurePackageFiles(t, NewSurePackageConfig(
		sourceRoot,
		errorHandlingMode,
		sourcePackagePath,
	))
}

func GenerateSurePackageFiles(t *testing.T, cfg *SurePackageGenConfig) {
	utils.PrintObject(cfg)

	for _, name := range utils.MustLs(cfg.SourceRoot) {
		if filepath.Ext(name) != ".go" {
			continue
		}
		if strings.HasSuffix(name, "_test.go") {
			continue
		}
		absPath := filepath.Join(cfg.SourceRoot, name)

		sureFunctionCodes := GenerateSureFunctions(t, cfg, absPath)

		if len(sureFunctionCodes) > 0 {
			shortErrorHandlingMode := strings.ToLower(string(cfg.ErrorHandlingMode))

			oldPackageName := syntaxgo.GetPkgName(absPath)

			newPackageName := zerotern.VV(cfg.NewPkgName, oldPackageName+"_"+shortErrorHandlingMode)

			ptx := utils.NewPTX()
			ptx.Println("package" + " " + newPackageName)
			ptx.Println("import(")
			ptx.Println(utils.SetDoubleQuotes(cfg.SourcePackagePath))
			ptx.Println(utils.SetDoubleQuotes(cfg.ErrorHandlingPkgPath))
			ptx.Println(")")
			ptx.Println(strings.Join(sureFunctionCodes, "\n"))

			newName := strings.Replace(name, ".go", "_"+shortErrorHandlingMode+".go", 1)
			newPath := filepath.Join(cfg.OutputRoot, newPackageName, newName)

			newCode, _ := formatgo.FormatCode(ptx.String())

			utils.MustWriteIntoPath(newPath, newCode)
		}
	}
}

func GenerateSureFunctions(t *testing.T, cfg *SurePackageGenConfig, absPath string) []string {
	utils.PrintObject(cfg)

	sourceCode := done.VAE(os.ReadFile(absPath)).Done()
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV1(sourceCode))

	packageName := astBundle.GetPackageName()

	astFile, _ := astBundle.GetBundle()

	astFunctions := syntaxgo_search.ExtractFunctions(astFile)

	var functionCodes []string
	for _, astFunc := range astFunctions {
		if astFunc.Recv != nil {
			if len(astFunc.Recv.List) > 0 && len(astFunc.Recv.List[0].Names) > 0 {
				t.Log(astFunc.Recv.List[0].Names[0].Name, astFunc.Name.Name)
			}
			continue
		}
		t.Log(astFunc.Name.Name)
		if !utils.C0IsUppercase(astFunc.Name.Name) {
			continue
		}
		results := parseFunctionReturnFields(sourceCode, astFunc)
		t.Log(utils.Neat2json(results))

		functionCode := makeFunctionCode(sourceCode, packageName, astFunc, results, cfg.ErrorHandlingMode, cfg.HandlerFuncReference)
		t.Log(functionCode)

		functionCodes = append(functionCodes, functionCode)
	}
	return functionCodes
}

func makeFunctionCode(
	sourceCode []byte,
	packageName string,
	astFunc *ast.FuncDecl,
	results []*resultType,
	errorHandlingMode sure.ErrorHandlingMode,
	sureHandlerFuncReference string,
) string {
	var newFunctionCode = "func " + astFunc.Name.Name
	if astFunc.Type.TypeParams != nil {
		newFunctionCode += syntaxgo_astnode.GetText(sourceCode, astFunc.Type.TypeParams)
	}

	genericTypeParams := syntaxgo_aktnorm.GetGenericTypeParamsMap(astFunc.Type.TypeParams)

	newFunctionCode += "("
	var argumentNames []string
	if astFunc.Type.Params != nil && len(astFunc.Type.Params.List) > 0 {
		var args []string
		for _, param := range astFunc.Type.Params.List {
			argType := syntaxgo_astnode.GetText(sourceCode, param.Type)
			argType = resolveFullExportedType(packageName, genericTypeParams, argType)

			if len(param.Names) == 0 {
				argName := "arg" + strconv.Itoa(len(args))

				args = append(args, argName+" "+argType)
				argumentNames = append(argumentNames, argName)
			} else {
				argNames := make([]string, 0, len(param.Names))

				for _, name := range param.Names {
					// 检查参数是否是 "..."
					if _, variadic := param.Type.(*ast.Ellipsis); variadic {
						argNames = append(argNames, name.Name+" ...")
					} else {
						argNames = append(argNames, name.Name)
					}
				}

				args = append(args, strings.Join(argNames, ",")+" "+argType)
				argumentNames = append(argumentNames, argNames...)
			}
		}
		newFunctionCode += strings.Join(args, ",")
	}
	newFunctionCode += ")"

	var anonymousReturn = false
	for _, res := range results {
		if res.IsAnonymous {
			anonymousReturn = true
			break
		}
	}

	var containsErrorReturn = false
	if len(results) > 0 {
		var resultNames = make([]string, 0, len(results))
		for _, res := range results {
			if res.Type == "error" {
				containsErrorReturn = true
				continue
			}
			if res.IsAnonymous {
				resultNames = append(resultNames, resolveFullExportedType(packageName, genericTypeParams, res.Type))
			} else {
				resultNames = append(resultNames, res.Name+" "+resolveFullExportedType(packageName, genericTypeParams, res.Type))
			}
		}
		if len(resultNames) > 0 {
			if len(resultNames) < 2 && anonymousReturn {
				newFunctionCode += " " + resultNames[0]
			} else {
				newFunctionCode += " " + "(" + strings.Join(resultNames, ",") + ")"
			}
		}
	}
	newFunctionCode += " " + "{" + "\n"
	if len(results) > 0 {
		var rets = make([]string, 0, len(results))
		for _, res := range results {
			rets = append(rets, res.Name)
		}
		newFunctionCode += strings.Join(rets, ",")
		if anonymousReturn || containsErrorReturn {
			newFunctionCode += ":="
		} else {
			newFunctionCode += "="
		}
	}
	newFunctionCode += packageName + "." + astFunc.Name.Name
	if astFunc.Type.TypeParams != nil {
		newFunctionCode += "["
		var genericTypeNames []string
		for _, xts := range astFunc.Type.TypeParams.List {
			for _, xtx := range xts.Names {
				genericTypeNames = append(genericTypeNames, xtx.Name)
			}
		}
		newFunctionCode += strings.Join(genericTypeNames, ",")
		newFunctionCode += "]"
	}
	newFunctionCode += "(" + strings.Join(argumentNames, ",") + ")" + "\n"
	if len(results) > 0 {
		for _, x := range results {
			if x.Type == "error" {
				newFunctionCode += sureHandlerFuncReference + "." + string(errorHandlingMode) + "(" + x.Name + ")" + "\n"
			}
		}
	}
	if len(results) > 0 {
		var rets = make([]string, 0, len(results))
		for _, res := range results {
			if res.Type == "error" {
				continue
			}
			rets = append(rets, res.Name)
		}
		if len(rets) > 0 {
			newFunctionCode += "return" + " " + strings.Join(rets, ",") + "\n"
		}
	}
	newFunctionCode += "}" + "\n"
	return newFunctionCode
}

func resolveFullExportedType(packageName string, genericTypeParams map[string]ast.Expr, resType string) string {
	if utils.C0IsUppercase(resType) {
		classType := resType
		if _, ok := genericTypeParams[classType]; ok {
			return resType
		}
		return packageName + "." + classType
	}
	if resType[0] == '*' {
		classType := resType[1:]
		if _, ok := genericTypeParams[classType]; ok {
			return resType
		}
		if utils.C0IsUppercase(classType) {
			return "*" + packageName + "." + classType
		}
	}
	return resType
}

type resultType struct {
	Name        string
	Type        string
	IsAnonymous bool
}

func parseFunctionReturnFields(source []byte, astFunction *ast.FuncDecl) []*resultType {
	var results []*resultType
	if astFunction.Type.Results == nil || len(astFunction.Type.Results.List) == 0 {
		results = make([]*resultType, 0)
	} else {
		var errNum int
		for _, elem := range astFunction.Type.Results.List {
			resType := syntaxgo_astnode.GetText(source, elem.Type)
			eIs := utils.Boolean(resType == "error")
			if len(elem.Names) == 0 {
				resName := tern.BFF(eIs, func() string {
					return tern.BVV(errNum == 0, "err", "err"+strconv.Itoa(errNum))
				}, func() string {
					return "res" + strconv.Itoa(len(results))
				})

				results = append(results, &resultType{
					Name:        resName,
					Type:        resType,
					IsAnonymous: true,
				})
				if eIs {
					errNum++
				}
			} else {
				for _, name := range elem.Names {
					results = append(results, &resultType{
						Name:        name.Name,
						Type:        resType,
						IsAnonymous: false,
					})
					if eIs {
						errNum++
					}
				}
			}
		}
	}
	return results
}
