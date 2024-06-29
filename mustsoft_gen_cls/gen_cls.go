package mustsoft_gen_cls

import (
	"go/ast"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/yyle88/done"
	"github.com/yyle88/formatgo"
	"github.com/yyle88/mustdone"
	"github.com/yyle88/mustdone/internal/utils"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_astfieldsflat"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type Config struct {
	GenParam      *GenParam
	PkgName       string
	ImportOptions *syntaxgo_ast.PackageImportOptions
	SrcPath       string
}

func Gen(cfg *Config, objects ...interface{}) {
	utils.StringOK(cfg.GenParam.SrcRoot)
	utils.StringOK(cfg.PkgName)
	utils.StringOK(cfg.SrcPath)

	ptx := utils.NewPTX()
	ptx.Println("package", cfg.PkgName)

	for _, object := range objects {
		ptx.Println(GenerateFlexibleClassCode(cfg.GenParam, object))
	}

	var importOptions *syntaxgo_ast.PackageImportOptions
	if cfg.ImportOptions != nil {
		importOptions = cfg.ImportOptions
	} else {
		importOptions = &syntaxgo_ast.PackageImportOptions{}
	}
	if cfg.GenParam.MustSoftCallableNode != "" {
		zaplog.LOG.Debug("use_new_must_soft_node", zap.String("node", cfg.GenParam.MustSoftCallableNode))
	} else { //表示使用的默认的 Must 和 Soft 函数，就说明你是需要引用这个包，补上有利于format代码
		importOptions.SetObject(syntaxgo_reflect.GetObject[mustdone.FlexibleEnum]())
	}

	//把需要 import 的包路径设置到代码里
	source := syntaxgo_ast.AddImports(ptx.Bytes(), importOptions)
	//统计 format 代码的时间
	startTime := time.Now()
	//执行 format 时，要确保它不再去找 imports 需要引用的包，否则就会比较耗时，当你发现这里很耗时时就可以顺着这个思路排查
	newSource := done.VAE(formatgo.FormatBytes(source)).Nice()
	//把格式化后的代码写到对应的文件路径里
	duration := time.Since(startTime)
	zaplog.LOG.Debug("gen", zap.Duration("format_cost_duration", duration))
	done.Done(utils.WriteFile(cfg.SrcPath, newSource))
	zaplog.LOG.Debug("gen_success")
}

func GenerateFlexibleClassCode(cfg *GenParam, object interface{}) string {
	ptx := utils.NewPTX()
	for _, flexibleEnum := range cfg.GetFlexibleEnums() {
		ptx.Println(GenerateFlexibleClassOnce(cfg, object, flexibleEnum))
	}
	return ptx.String()
}

func GenerateFlexibleClassOnce(cfg *GenParam, object interface{}, flexibleEnum mustdone.FlexibleEnum) string {
	objectType := reflect.TypeOf(object)
	zaplog.LOG.Debug(utils.StringOK(objectType.Name()))
	zaplog.LOG.Debug(utils.StringOK(objectType.String()))
	zaplog.LOG.Debug(utils.StringOK(objectType.PkgPath()))

	utils.RootMustIsExist(cfg.SrcRoot)

	utils.BooleanOK(flexibleEnum == mustdone.MUST || flexibleEnum == mustdone.SOFT)

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

		source := done.VAE(os.ReadFile(path)).Done()

		astFile := done.VCE(syntaxgo_ast.NewAstFromSource(source)).Nice()
		astFns := syntaxgo_ast.GetFunctions(astFile)
		mebFunctions := syntaxgo_ast.GetFunctionsXRecvName(astFns, objectType.Name(), true)
		if len(mebFunctions) == 0 {
			continue
		}

		astTuples = append(astTuples, &srcFnsTuple{
			srcCode: source,
			methods: mebFunctions,
		})
	}

	if cfg.SubClassRecvName == "" {
		cfg.SubClassRecvName = utils.SOrX(astTuples.GetRecvName(), "T")
	}

	subClassName := cfg.makeClassName(reflect.TypeOf(object), flexibleEnum)

	ptx := utils.NewPTX()
	ptx.Println(`type ` + subClassName + ` struct{` + cfg.SubClassRecvName + ` *` + objectType.Name() + `}`)
	ptx.Println(`
		func(` + cfg.SubClassRecvName + ` *` + objectType.Name() + `) ` + string(flexibleEnum) + `() * ` + subClassName + `{
		return & ` + subClassName + `{` + cfg.SubClassRecvName + `:` + cfg.SubClassRecvName + `}
	}`)

	for _, oneTmp := range astTuples {
		source := oneTmp.srcCode
		mebFunctions := oneTmp.methods
		for _, mebFunc := range mebFunctions {
			mebFuncName := syntaxgo_ast.GetNodeCode(source, mebFunc.Name)
			if utils.In(mebFuncName, []string{string(mustdone.MUST), string(mustdone.SOFT)}) {
				continue
			}

			var params = make(syntaxgo_astfieldsflat.NameTypeElements, 0)
			if mebFunc.Type != nil && mebFunc.Type.Params != nil {
				params = syntaxgo_astfieldsflat.GetSimpleArgElements(mebFunc.Type.Params.List, source)
			}
			var results = make(syntaxgo_astfieldsflat.NameTypeElements, 0)
			if mebFunc.Type != nil && mebFunc.Type.Results != nil {
				results = syntaxgo_astfieldsflat.GetSimpleResElements(mebFunc.Type.Results.List, source)
			}

			for _, elem := range results {
				zaplog.LOG.Debug("elem", zap.String("name", elem.Name), zap.String("kind", elem.Kind))
			}

			var okxResElems = make(syntaxgo_astfieldsflat.NameTypeElements, 0)
			var erxResElems = make(syntaxgo_astfieldsflat.NameTypeElements, 0)
			for _, result := range results {
				if syntaxgo_ast.GetNodeCode(source, result.Type) == "error" {
					erxResElems = append(erxResElems, result)
				} else {
					okxResElems = append(okxResElems, result)
				}
			}

			var erxHandleStmts []string
			for _, erxName := range erxResElems.GetFunctionParamsStats() {
				var callableNode string
				if cfg.MustSoftCallableNode != "" {
					callableNode = cfg.MustSoftCallableNode
				} else {
					callableNode = mustdone.GetPkgName()
				}
				erxHandleStmts = append(erxHandleStmts, callableNode+"."+string(flexibleEnum)+"("+erxName+")")
			}

			ptx.Println(`func (T *` + subClassName + `) ` + mebFuncName + `(` +
				params.GetNamesKindsStats().MergeParts() +
				`)` + `(` +
				okxResElems.GetNamesKindsStats().MergeParts() +
				`) {`)

			runFuncLine := `T.` + cfg.SubClassRecvName + `.` + mebFuncName + `(` + params.GetFunctionParamsStats().MergeParts() + `)`
			if len(results) > 0 {
				if len(okxResElems) == len(results) {
					ptx.Println(results.GetFunctionParamsStats().MergeParts() + "=" + runFuncLine)
				} else {
					ptx.Println(results.GetFunctionParamsStats().MergeParts() + ":=" + runFuncLine)
				}
			} else {
				ptx.Println(runFuncLine)
			}

			if len(erxHandleStmts) > 0 {
				ptx.Println(strings.Join(erxHandleStmts, "\n"))
			}

			if len(okxResElems) > 0 {
				ptx.Println("return" + " " + okxResElems.GetFunctionParamsStats().MergeParts())
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

func (vs srcFnsTuples) GetRecvName() string {
	for _, oneTmp := range vs {
		for _, mebFunction := range oneTmp.methods {
			if mebFunction.Recv == nil {
				continue
			}
			if len(mebFunction.Recv.List) == 0 {
				continue
			}
			if len(mebFunction.Recv.List[0].Names) == 0 {
				continue
			}
			name := mebFunction.Recv.List[0].Names[0].Name
			if name != "" {
				return name
			}
		}
	}
	return ""
}
