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
	"github.com/yyle88/sure/internal/utils"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_astfieldsflat"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type Param struct {
	object any
	opStub string
}

func NewParam(object any, opStub string) *Param {
	return &Param{
		object: object,
		opStub: opStub,
	}
}

type Config struct {
	SrcRoot       string
	TargetPkgName string
	ImportOptions *syntaxgo_ast.PackageImportOptions
	TargetSrcPath string
}

func Gen(cfg *Config, params ...*Param) {
	utils.StringOK(cfg.SrcRoot)
	utils.StringOK(cfg.TargetPkgName)
	utils.StringOK(cfg.TargetSrcPath)

	ptx := utils.NewPTX()
	ptx.Println("package", cfg.TargetPkgName)

	for _, param := range params {
		ptx.Println(GenerateStubFunctions(cfg, param))
	}

	var importOptions *syntaxgo_ast.PackageImportOptions
	if cfg.ImportOptions != nil {
		importOptions = cfg.ImportOptions
	} else {
		importOptions = &syntaxgo_ast.PackageImportOptions{}
	}

	for _, param := range params {
		importOptions.SetPkgPath(syntaxgo_reflect.GetPkgPathV3(param.object))
	}

	//把需要 import 的包路径设置到代码里
	source := syntaxgo_ast.AddImports(ptx.Bytes(), importOptions)

	zaplog.LOG.Debug(string(source))

	//统计 format 代码的时间
	startTime := time.Now()
	//执行 format 时，要确保它不再去找 imports 需要引用的包，否则就会比较耗时，当你发现这里很耗时时就可以顺着这个思路排查
	newSource := done.VAE(formatgo.FormatBytes(source)).Nice()

	//把格式化后的代码写到对应的文件路径里
	duration := time.Since(startTime)
	zaplog.LOG.Debug("gen", zap.Duration("format_cost_duration", duration))
	if cfg.TargetSrcPath != "" {
		done.Done(utils.WriteFile(cfg.TargetSrcPath, newSource))
	} else {
		fmt.Println(newSource)
	}
	zaplog.LOG.Debug("gen_success")
}

func GenerateStubFunctions(cfg *Config, param *Param) string {
	objectType := syntaxgo_reflect.GetTypeV3(param.object)
	zaplog.LOG.Debug(utils.StringOK(objectType.Name()))
	zaplog.LOG.Debug(utils.StringOK(objectType.String()))
	zaplog.LOG.Debug(utils.StringOK(objectType.PkgPath()))

	utils.MustRoot(cfg.SrcRoot)

	var astTuples = make(srcFnsTuples, 0)
	for _, subInfo := range done.VAE(os.ReadDir(cfg.SrcRoot)).Done() {
		if subInfo.IsDir() {
			continue
		}
		if !(filepath.Ext(subInfo.Name()) == ".go") {
			continue
		}
		path := filepath.Join(cfg.SrcRoot, subInfo.Name())
		zaplog.LOG.Debug(path)

		srcCode := done.VAE(os.ReadFile(path)).Done()

		astFile := done.VCE(syntaxgo_ast.NewAstFromSource(srcCode)).Nice()
		astFcXs := syntaxgo_ast.GetFunctions(astFile)
		methods := syntaxgo_ast.GetFunctionsXRecvName(astFcXs, objectType.Name(), true)
		if len(methods) == 0 {
			continue
		}

		astTuples = append(astTuples, &srcFnsTuple{
			srcCode: srcCode,
			methods: methods,
		})
	}

	ptx := utils.NewPTX()

	for _, oneTmp := range astTuples {
		srcCode := oneTmp.srcCode
		methods := oneTmp.methods
		for _, mebFunc := range methods {
			mebFuncName := syntaxgo_ast.GetNodeCode(srcCode, mebFunc.Name)

			var params = make(syntaxgo_astfieldsflat.NameTypeElements, 0)
			if mebFunc.Type != nil && mebFunc.Type.Params != nil {
				params = syntaxgo_astfieldsflat.GetSimpleArgElements(mebFunc.Type.Params.List, srcCode)

				for _, elem := range params {
					elem.SetPkgUsage(syntaxgo_reflect.GetPkgNameV3(param.object), make(map[string]int))
				}
			}
			var results = make(syntaxgo_astfieldsflat.NameTypeElements, 0)
			if mebFunc.Type != nil && mebFunc.Type.Results != nil {
				results = syntaxgo_astfieldsflat.GetSimpleResElements(mebFunc.Type.Results.List, srcCode)

				for _, elem := range results {
					elem.SetPkgUsage(syntaxgo_reflect.GetPkgNameV3(param.object), make(map[string]int))
				}
			}

			for _, elem := range results {
				zaplog.LOG.Debug("elem", zap.String("name", elem.Name), zap.String("kind", elem.Kind))
			}

			ptx.Println(`func ` + mebFuncName + `(` +
				params.GetNamesKindsStats().MergeParts() +
				`)` + `(` +
				strings.Join(results.Kinds(), ",") +
				`) {`)

			runFuncLine := param.opStub + `.` + mebFuncName + `(` + params.GetFunctionParamsStats().MergeParts() + `)`
			if len(results) > 0 {
				ptx.Println("return " + runFuncLine)
			} else {
				ptx.Println(runFuncLine)
			}

			ptx.Println("}")
		}
	}
	res := ptx.String()
	zaplog.LOG.Debug(res)
	return res
}

type srcFnsTuple struct {
	srcCode []byte
	methods []*ast.FuncDecl
}

type srcFnsTuples []*srcFnsTuple
