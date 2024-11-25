package sure_pkg_gen

import (
	"fmt"
	"go/ast"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/yyle88/done"
	"github.com/yyle88/formatgo"
	"github.com/yyle88/sure"
	"github.com/yyle88/sure/internal/utils"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_astvtnorm"
	"github.com/yyle88/tern"
	"github.com/yyle88/tern/zerotern"
)

func Gen(t *testing.T, pkgRoot string, sureEnum sure.SureEnum, pkgPath string) {
	GenerateSurePackage(t, NewConfig(pkgRoot, sureEnum, pkgPath))
}

func GenerateSurePackage(t *testing.T, cfg *Config) {
	fmt.Println(cfg.PkgRoot, cfg.GenRoot)

	for _, name := range utils.MustLs(cfg.PkgRoot) {
		if filepath.Ext(name) != ".go" {
			continue
		}
		if strings.HasSuffix(name, "_test.go") {
			continue
		}
		absPath := filepath.Join(cfg.PkgRoot, name)
		srcData := done.VAE(os.ReadFile(absPath)).Done()

		astFile, err := syntaxgo_ast.NewAstXFilepath(absPath)
		done.Done(err)

		packageName := astFile.Name.Name

		astFcXs := syntaxgo_ast.GetFunctions(astFile)

		var sliceFuncCodes []string
		for _, astFunc := range astFcXs {
			if astFunc.Recv != nil {
				if len(astFunc.Recv.List) > 0 && len(astFunc.Recv.List[0].Names) > 0 {
					t.Log(astFunc.Recv.List[0].Names[0].Name, astFunc.Name.Name)
				}
				continue
			}
			t.Log(astFunc.Name.Name)
			if !utils.C0IsUPPER(astFunc.Name.Name) {
				continue
			}
			results, anonymous := parseResFields(srcData, astFunc)
			t.Log(utils.Neat(results))

			sFuncCode := newFuncCode(srcData, packageName, astFunc, results, anonymous, cfg.SureEnum, cfg.SureUseNode)
			t.Log(sFuncCode)

			sliceFuncCodes = append(sliceFuncCodes, sFuncCode)
		}

		if len(sliceFuncCodes) > 0 {
			shortSureName := strings.ToLower(string(cfg.SureEnum))
			newPackageName := zerotern.VV(cfg.NewPkgName, packageName+"_"+shortSureName)

			ptx := utils.NewPTX()
			ptx.Println("package" + " " + newPackageName)
			ptx.Println("import(")
			ptx.Println(utils.SetDoubleQuotes(cfg.PkgPath))
			ptx.Println(utils.SetDoubleQuotes(cfg.SurePkgPath))
			ptx.Println(")")
			ptx.Println(strings.Join(sliceFuncCodes, "\n"))

			newName := strings.Replace(name, ".go", "_"+shortSureName+".go", 1)
			newPath := filepath.Join(cfg.GenRoot, newPackageName, newName)

			newCode, _ := formatgo.FormatCode(ptx.String())

			utils.MustWriteIntoPath(newPath, newCode)
		}
	}
}

func newFuncCode(srcData []byte, packageName string, astFunc *ast.FuncDecl, results []*retType, anonymous bool, sureEnum sure.SureEnum, sureUseNode string) string {
	var res = "func " + astFunc.Name.Name
	if astFunc.Type.TypeParams != nil {
		res += syntaxgo_ast.GetNodeCode(srcData, astFunc.Type.TypeParams)
	}
	res += "("
	var argList []string
	if astFunc.Type.Params != nil && len(astFunc.Type.Params.List) > 0 {
		var args []string
		for _, param := range astFunc.Type.Params.List {
			if len(param.Names) == 0 {
				argName := "arg" + strconv.Itoa(len(args))
				args = append(args, argName+" "+syntaxgo_ast.GetNodeCode(srcData, param.Type))
				argList = append(argList, argName)
			} else {
				args = append(args, syntaxgo_ast.GetNodeCode(srcData, param))
				for _, name := range param.Names {
					// 检查参数是否是 "..."
					if _, variadic := param.Type.(*ast.Ellipsis); variadic {
						argList = append(argList, name.Name+" ...")
					} else {
						argList = append(argList, name.Name)
					}
				}
			}
		}
		res += strings.Join(args, ",")
	}
	res += ")"

	genericsMap := syntaxgo_astvtnorm.CountGenericsMap(astFunc.Type.TypeParams)

	var isReturnErrors = false
	{
		var rets = make([]string, 0, len(results))
		if anonymous {
			for _, res := range results {
				if res.Type == "error" {
					isReturnErrors = true
					continue
				}
				rets = append(rets, cvtAZType(packageName, genericsMap, res))
			}
		} else {
			for _, res := range results {
				if res.Type == "error" {
					isReturnErrors = true
					continue
				}
				rets = append(rets, res.Name+" "+cvtAZType(packageName, genericsMap, res))
			}
		}
		if len(rets) > 0 {
			if len(rets) < 2 && anonymous {
				res += " " + rets[0]
			} else {
				res += " " + "(" + strings.Join(rets, ",") + ")"
			}
		}
	}
	res += " " + "{" + "\n"
	if len(results) > 0 {
		var rets = make([]string, 0, len(results))
		for _, res := range results {
			rets = append(rets, res.Name)
		}
		res += strings.Join(rets, ",")
		if anonymous || isReturnErrors {
			res += ":="
		} else {
			res += "="
		}
	}
	res += packageName + "." + astFunc.Name.Name
	if astFunc.Type.TypeParams != nil {
		res += "["
		var genericTypeNames []string
		for _, xts := range astFunc.Type.TypeParams.List {
			for _, xtx := range xts.Names {
				genericTypeNames = append(genericTypeNames, xtx.Name)
			}
		}
		res += strings.Join(genericTypeNames, ",")
		res += "]"
	}
	res += "(" + strings.Join(argList, ",") + ")" + "\n"
	if len(results) > 0 {
		for _, x := range results {
			if x.Type == "error" {
				res += sureUseNode + "." + string(sureEnum) + "(" + x.Name + ")" + "\n"
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
			res += "return" + " " + strings.Join(rets, ",") + "\n"
		}
	}
	res += "}" + "\n"
	return res
}

func cvtAZType(packageName string, genericsMap map[string]int, res *retType) string {
	if utils.C0IsUPPER(res.Type) {
		classType := res.Type
		if _, ok := genericsMap[classType]; ok {
			return res.Type
		}
		return packageName + "." + classType
	}
	if res.Type[0] == '*' {
		classType := res.Type[1:]
		if _, ok := genericsMap[classType]; ok {
			return res.Type
		}
		if utils.C0IsUPPER(classType) {
			return "*" + packageName + "." + classType
		}
	}
	return res.Type
}

type retType struct {
	Name string
	Type string
}

func parseResFields(source []byte, astFunction *ast.FuncDecl) ([]*retType, bool) {
	var results []*retType
	var anonymous = true
	if astFunction.Type.Results == nil || len(astFunction.Type.Results.List) == 0 {
		results = make([]*retType, 0)
	} else {
		var errNum int
		for _, x := range astFunction.Type.Results.List {
			resType := syntaxgo_ast.GetNodeCode(source, x.Type)
			eIs := utils.Boolean(resType == "error")
			if len(x.Names) == 0 {
				resName := tern.BFF(eIs, func() string {
					return tern.BVV(errNum == 0, "err", "err"+strconv.Itoa(errNum))
				}, func() string {
					return "res" + strconv.Itoa(len(results))
				})

				results = append(results, &retType{
					Name: resName,
					Type: resType,
				})
				if eIs {
					errNum++
				}
			} else {
				anonymous = false
				for _, name := range x.Names {
					results = append(results, &retType{
						Name: name.Name,
						Type: resType,
					})
					if eIs {
						errNum++
					}
				}
			}
		}
	}
	return results, anonymous
}
