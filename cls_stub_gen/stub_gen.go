package cls_stub_gen

import (
	"fmt"
	"go/ast"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/yyle88/done"
	"github.com/yyle88/formatgo"
	"github.com/yyle88/must"
	"github.com/yyle88/sure/internal/utils"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_astnode"
	"github.com/yyle88/syntaxgo/syntaxgo_astnorm"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
	"github.com/yyle88/syntaxgo/syntaxgo_search"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type StubParam struct {
	object             any
	operationNamespace string
}

func NewStubParam(object any, operationNamespace string) *StubParam {
	return &StubParam{
		object:             object,
		operationNamespace: operationNamespace,
	}
}

type StubGenConfig struct {
	SourceRootPath    string
	TargetPackageName string
	ImportOptions     *syntaxgo_ast.PackageImportOptions
	OutputPath        string
	AllowFileCreation bool //当目标文件不存在时，能否能新建文件，假如设置为不能，就必须得找到文件才能写内容
}

func GenerateStubs(config *StubGenConfig, stubParams ...*StubParam) {
	utils.PrintObject(config)

	must.Nice(config.SourceRootPath)
	must.Nice(config.TargetPackageName)
	must.Nice(config.OutputPath)

	ptx := utils.NewPTX()
	ptx.Println("package", config.TargetPackageName)

	for _, param := range stubParams {
		ptx.Println(GenerateStubMethods(config, param))
	}

	var importOptions *syntaxgo_ast.PackageImportOptions
	if config.ImportOptions != nil {
		importOptions = config.ImportOptions
	} else {
		importOptions = &syntaxgo_ast.PackageImportOptions{}
	}

	for _, param := range stubParams {
		importOptions.SetPkgPath(syntaxgo_reflect.GetPkgPathV3(param.object))
	}

	//把需要 import 的包路径设置到代码里
	source := importOptions.InjectImports(ptx.Bytes())

	zaplog.LOG.Debug(string(source))

	//统计 format 代码的时间
	sinceTime := time.Now()
	//执行 format 时，要确保它不再去找 imports 需要引用的包，否则就会比较耗时，当你发现这里很耗时时就可以顺着这个思路排查
	newSource := done.VAE(formatgo.FormatBytes(source)).Nice()

	//把格式化后的代码写到对应的文件路径里
	duration := time.Since(sinceTime)
	zaplog.LOG.Debug("gen", zap.Duration("format_cost_duration", duration))
	if config.OutputPath != "" {
		if config.AllowFileCreation { //当确定目标不是必须存在时，就不检查目标文件，但这有可能导致配置错误而写到错误的地方
			zaplog.LOG.Warn("write to target src path", zap.String("path", config.OutputPath))
		} else { //推荐还是首先自己把文件建好，让程序能找到这个文件，再去重写它的内容
			utils.MustFile(config.OutputPath) //规定目标是必须存在的
		}
		done.Done(os.WriteFile(config.OutputPath, newSource, 0644))
	} else {
		fmt.Println(string(newSource))
	}
	zaplog.LOG.Debug("gen_success")
}

func GenerateStubMethods(cfg *StubGenConfig, stubParam *StubParam) string {
	utils.PrintObject(cfg)

	objectType := syntaxgo_reflect.GetTypeV3(stubParam.object)
	zaplog.LOG.Debug(must.Nice(objectType.Name()))
	zaplog.LOG.Debug(must.Nice(objectType.String()))
	zaplog.LOG.Debug(must.Nice(objectType.PkgPath()))

	utils.MustRoot(cfg.SourceRootPath)

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

	ptx := utils.NewPTX()

	for _, sourceMethodsTuple := range sourceMethodsTuples {
		sourceCode := sourceMethodsTuple.SourceCode
		methodList := sourceMethodsTuple.MethodList
		for _, methodFunction := range methodList {
			methodName := syntaxgo_astnode.GetText(sourceCode, methodFunction.Name)

			var methodParameters = make(syntaxgo_astnorm.NameTypeElements, 0)
			if methodFunction.Type != nil && methodFunction.Type.Params != nil {
				methodParameters = syntaxgo_astnorm.GetSimpleArgElements(methodFunction.Type.Params.List, sourceCode)

				for _, elem := range methodParameters {
					elem.AdjustTypeWithPackage(syntaxgo_reflect.GetPkgNameV3(stubParam.object), make(map[string]ast.Expr))
				}
			}
			var returnValues = make(syntaxgo_astnorm.NameTypeElements, 0)
			if methodFunction.Type != nil && methodFunction.Type.Results != nil {
				returnValues = syntaxgo_astnorm.GetSimpleResElements(methodFunction.Type.Results.List, sourceCode)

				for _, elem := range returnValues {
					elem.AdjustTypeWithPackage(syntaxgo_reflect.GetPkgNameV3(stubParam.object), make(map[string]ast.Expr))
				}
			}

			for _, elem := range returnValues {
				zaplog.LOG.Debug("elem", zap.String("name", elem.Name), zap.String("kind", elem.Kind))
			}

			ptx.Println(`func ` + methodName + `(` +
				methodParameters.FormatNamesWithKinds().MergeParts() +
				`)` + `(` +
				strings.Join(returnValues.Kinds(), ",") +
				`) {`)

			methodInvocationStatement := stubParam.operationNamespace + `.` + methodName + `(` + methodParameters.GenerateFunctionParams().MergeParts() + `)`
			if len(returnValues) > 0 {
				ptx.Println("return " + methodInvocationStatement)
			} else {
				ptx.Println(methodInvocationStatement)
			}

			ptx.Println("}")
		}
	}
	res := ptx.String()
	zaplog.LOG.Debug(res)
	return res
}
