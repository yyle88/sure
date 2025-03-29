package sure_cls_gen

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/yyle88/done"
	"github.com/yyle88/formatgo"
	"github.com/yyle88/must"
	"github.com/yyle88/sure"
	"github.com/yyle88/sure/internal/utils"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_astnode"
	"github.com/yyle88/syntaxgo/syntaxgo_astnorm"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
	"github.com/yyle88/syntaxgo/syntaxgo_search"
	"github.com/yyle88/tern/zerotern"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type ClassGenConfig struct {
	ClassGenOptions *ClassGenOptions
	PackageName     string
	ImportOptions   *syntaxgo_ast.PackageImportOptions
	OutputPath      string
}

func GenerateClasses(cfg *ClassGenConfig, objects ...interface{}) {
	utils.PrintObject(cfg)

	must.Nice(cfg.ClassGenOptions.SourceRootPath)
	must.Nice(cfg.PackageName)
	must.Nice(cfg.OutputPath)

	ptx := utils.NewPTX()
	ptx.Println("package", cfg.PackageName)

	for _, object := range objects {
		ptx.Println(GenerateClassCodeWithErrorHandling(cfg.ClassGenOptions, object))
	}

	var importOptions *syntaxgo_ast.PackageImportOptions
	if cfg.ImportOptions != nil {
		importOptions = cfg.ImportOptions
	} else {
		importOptions = &syntaxgo_ast.PackageImportOptions{}
	}
	if cfg.ClassGenOptions.ErrorHandlerFuncName != "" {
		zaplog.LOG.Debug("use_new_sure_node", zap.String("node", cfg.ClassGenOptions.ErrorHandlerFuncName))
	} else { //表示使用的默认的 Must 和 Soft 函数，就说明你是需要引用这个包，补上有利于format代码
		importOptions.SetInferredObject(syntaxgo_reflect.GetObject[sure.ErrorHandlingMode]())
	}

	//把需要 import 的包路径设置到代码里
	source := importOptions.InjectImports(ptx.Bytes())
	//统计 format 代码的时间
	sinceTime := time.Now()
	//执行 format 时，要确保它不再去找 imports 需要引用的包，否则就会比较耗时，当你发现这里很耗时时就可以顺着这个思路排查
	newSource := done.VAE(formatgo.FormatBytes(source)).Nice()
	//把格式化后的代码写到对应的文件路径里
	duration := time.Since(sinceTime)
	zaplog.LOG.Debug("gen", zap.Duration("format_cost_duration", duration))
	if cfg.OutputPath != "" {
		done.Done(utils.WriteFile(cfg.OutputPath, newSource))
	} else {
		fmt.Println(string(newSource))
	}
	zaplog.LOG.Debug("gen_success")
}

func GenerateClassCodeWithErrorHandling(cfg *ClassGenOptions, object interface{}) string {
	utils.PrintObject(cfg)

	ptx := utils.NewPTX()
	for _, errorHandlingMode := range cfg.GetErrorHandlingModes() {
		ptx.Println(GenerateClassWithErrorHandlingMode(cfg, object, errorHandlingMode))
	}
	return ptx.String()
}

func GenerateClassWithErrorHandlingMode(cfg *ClassGenOptions, object interface{}, errorHandlingMode sure.ErrorHandlingMode) string {
	utils.PrintObject(cfg)

	objectType := reflect.TypeOf(object)
	zaplog.LOG.Debug(must.Nice(objectType.Name()))
	zaplog.LOG.Debug(must.Nice(objectType.String()))
	zaplog.LOG.Debug(must.Nice(objectType.PkgPath()))

	utils.MustRoot(cfg.SourceRootPath)

	if len(cfg.ErrorHandlingModes) == 0 { //当不填的时候就只能是默认的这两个枚举，而当填的时候允许开发者自定义别的
		must.True(errorHandlingMode == sure.MUST || errorHandlingMode == sure.SOFT)
	}

	var sourceMethodsTuples = make(utils.SourceMethodsTuples, 0)
	for _, fileInfo := range done.VAE(os.ReadDir(cfg.SourceRootPath)).Done() {
		if fileInfo.IsDir() {
			continue
		}
		if filepath.Ext(fileInfo.Name()) != ".go" {
			continue
		}
		path := filepath.Join(cfg.SourceRootPath, fileInfo.Name())
		zaplog.LOG.Debug(path)

		sourceCode := done.VAE(os.ReadFile(path)).Done()

		astBundle := done.VCE(syntaxgo_ast.NewAstBundleV1(sourceCode)).Nice()

		astFile, _ := astBundle.GetBundle()

		extractedFunctions := syntaxgo_search.ExtractFunctions(astFile)
		methodList := syntaxgo_search.ExtractFunctionsByReceiverName(extractedFunctions, objectType.Name(), true)
		if len(methodList) == 0 {
			continue
		}

		sourceMethodsTuples = append(sourceMethodsTuples, &utils.SourceMethodsTuple{
			SourceCode: sourceCode,
			MethodList: methodList,
		})
	}

	// when zero - set a new value
	zerotern.SetPF(&cfg.ReceiverVariableName, func() string {
		return zerotern.VV(sourceMethodsTuples.GetReceiverVariableName(), "T")
	})

	newClassName := cfg.GenerateNewClassName(reflect.TypeOf(object), errorHandlingMode)

	ptx := utils.NewPTX()
	ptx.Println(`type ` + newClassName + ` struct{` + cfg.ReceiverVariableName + ` *` + objectType.Name() + `}`)
	ptx.Println(`
		func(` + cfg.ReceiverVariableName + ` *` + objectType.Name() + `) ` + string(errorHandlingMode) + `() * ` + newClassName + `{
		return & ` + newClassName + `{` + cfg.ReceiverVariableName + `:` + cfg.ReceiverVariableName + `}
	}`)

	for _, sourceMethodsTuple := range sourceMethodsTuples {
		sourceCode := sourceMethodsTuple.SourceCode
		methodList := sourceMethodsTuple.MethodList
		for _, methodFunc := range methodList {
			methodName := syntaxgo_astnode.GetText(sourceCode, methodFunc.Name)
			if utils.In(methodName, []string{string(sure.MUST), string(sure.SOFT)}) {
				continue
			}
			if utils.In(sure.ErrorHandlingMode(methodName), cfg.ErrorHandlingModes) {
				continue
			}

			var params = make(syntaxgo_astnorm.NameTypeElements, 0)
			if methodFunc.Type != nil && methodFunc.Type.Params != nil {
				params = syntaxgo_astnorm.GetSimpleArgElements(methodFunc.Type.Params.List, sourceCode)
			}
			var results = make(syntaxgo_astnorm.NameTypeElements, 0)
			if methodFunc.Type != nil && methodFunc.Type.Results != nil {
				results = syntaxgo_astnorm.GetSimpleResElements(methodFunc.Type.Results.List, sourceCode)
			}

			for _, elem := range results {
				zaplog.LOG.Debug("elem", zap.String("name", elem.Name), zap.String("kind", elem.Kind))
			}

			var rightResultElements = make(syntaxgo_astnorm.NameTypeElements, 0)
			var errorResultElements = make(syntaxgo_astnorm.NameTypeElements, 0)
			for _, result := range results {
				if syntaxgo_astnode.GetText(sourceCode, result.Type) == "error" {
					errorResultElements = append(errorResultElements, result)
				} else {
					rightResultElements = append(rightResultElements, result)
				}
			}

			var errorHandlingStatements []string
			for _, errName := range errorResultElements.GenerateFunctionParams() {
				sureNode := zerotern.VF(cfg.ErrorHandlerFuncName, sure.GetPkgName)

				errorHandlingStatements = append(errorHandlingStatements, sureNode+"."+string(errorHandlingMode)+"("+errName+")")
			}

			ptx.Println(`func (T *` + newClassName + `) ` + methodName + `(` +
				params.FormatNamesWithKinds().MergeParts() +
				`)` + `(` +
				rightResultElements.FormatNamesWithKinds().MergeParts() +
				`) {`)

			methodInvocationStatement := `T.` + cfg.ReceiverVariableName + `.` + methodName + `(` + params.GenerateFunctionParams().MergeParts() + `)`
			if len(results) > 0 {
				if len(rightResultElements) == len(results) {
					ptx.Println(results.GenerateFunctionParams().MergeParts() + "=" + methodInvocationStatement)
				} else {
					ptx.Println(results.GenerateFunctionParams().MergeParts() + ":=" + methodInvocationStatement)
				}
			} else {
				ptx.Println(methodInvocationStatement)
			}

			if len(errorHandlingStatements) > 0 {
				ptx.Println(strings.Join(errorHandlingStatements, "\n"))
			}

			if len(rightResultElements) > 0 {
				ptx.Println("return" + " " + rightResultElements.GenerateFunctionParams().MergeParts())
			}

			ptx.Println("}")
		}
	}
	res := ptx.String()
	zaplog.LOG.Debug(res)
	return res
}
