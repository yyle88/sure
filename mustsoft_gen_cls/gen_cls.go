package mustsoft_gen_cls

import (
	"go/ast"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/yyle88/done"
	"github.com/yyle88/mustdone"
	"github.com/yyle88/mustdone/internal/utils"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_astfieldsflat"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func Gen(cfg *GenParam, object any, flexibleEnums ...mustdone.FlexibleEnum) string {
	ptx := utils.NewPTX()
	for _, flexibleEnum := range flexibleEnums {
		ptx.Println(GenerateFlexibleClassCode(cfg, object, flexibleEnum))
	}
	return ptx.String()
}

func GenerateFlexibleClassCode(cfg *GenParam, object any, flexibleEnum mustdone.FlexibleEnum) string {
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
		if !strings.HasSuffix(subInfo.Name(), ".go") {
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

	if cfg.OptRecvName == "" {
		cfg.OptRecvName = utils.SOrX(astTuples.GetRecvName(), "T")
	}

	subClassName := cfg.makeClassName(reflect.TypeOf(object), flexibleEnum)

	ptx := utils.NewPTX()
	ptx.Println(`type ` + subClassName + ` struct{` + cfg.OptRecvName + ` *` + objectType.Name() + `}`)
	ptx.Println(`
		func(` + cfg.OptRecvName + ` *` + objectType.Name() + `) ` + string(flexibleEnum) + `() * ` + subClassName + `{
		return & ` + subClassName + `{` + cfg.OptRecvName + `:` + cfg.OptRecvName + `}
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
			for _, erxElemName := range erxResElems.GetFunctionParamsStats() {
				className := cfg.FlexClass
				if className == "" {
					className = mustdone.GetPkgName()
				}
				erxHandleStmts = append(erxHandleStmts, className+"."+string(flexibleEnum)+"("+erxElemName+")")
			}

			ptx.Println(`func (T *` + subClassName + `) ` + mebFuncName + `(` +
				params.GetNamesKindsStats().MergeParts() +
				`)` + `(` +
				okxResElems.GetNamesKindsStats().MergeParts() +
				`) {`)

			runFuncLine := `T.` + cfg.OptRecvName + `.` + mebFuncName + `(` + params.GetFunctionParamsStats().MergeParts() + `)`
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
