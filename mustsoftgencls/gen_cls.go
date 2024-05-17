package mustsoftgencls

import (
	"go/ast"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/yyle88/done"
	"github.com/yyle88/mustdone"
	"github.com/yyle88/mustdone/internal/utils"
	"github.com/yyle88/zaplog"
	"gitlab.yyle.com/golang/uvyyle.git/utils_golang/utils_golang_ast"
	"gitlab.yyle.com/golang/uvyyle.git/utils_golang/utils_golang_ast/utils_golang_ast_fields"
	"go.uber.org/zap"
)

type GenParam struct {
	SourceCodeRootPath    string    //必填参数，你要解析的类型所在源代码的目录
	OptRecvName           string    //默认不填，你要解析的类型它的成员函数的recv的名称，比如 func (a *A)do() 就填写 a 就行
	SubClassNameStyleType StyleType //非必填参数，你要生成的新子类型的命名风格，有默认风格
	SubClassNamePartWords string    //非必填参数，你要生成的新子类型的名称片段，就是这个名称字符串中间，有部分可以自定义的内容
	FlexClass             string    //非必填参数，就是调用 FLEX 函数的调用者，你也可以实现自己的 flex 函数，默认用 flex 包的
}

func GenCodes(object any, cfg *GenParam, flexibleTypes ...mustdone.FlexibleHandlingType) string {
	ptx := utils.NewPTX()
	for _, flexibleType := range flexibleTypes {
		subClassName := makeNewTypeName(reflect.TypeOf(object), cfg, flexibleType)
		ptx.Println(GenCode(object, cfg, flexibleType, subClassName))
	}
	return ptx.String()
}

func GenCode(object any, cfg *GenParam, flexibleType mustdone.FlexibleHandlingType, subClassName string) string {
	objectType := reflect.TypeOf(object)
	zaplog.LOG.Debug(utils.StringOK(objectType.Name()))
	zaplog.LOG.Debug(utils.StringOK(objectType.String()))
	zaplog.LOG.Debug(utils.StringOK(objectType.PkgPath()))

	utils.RootMustIsExist(cfg.SourceCodeRootPath)

	utils.BooleanOK(flexibleType == mustdone.MUST || flexibleType == mustdone.SOFT)

	var astTuples = make(astSrcTmpTuples, 0)
	for _, subInfo := range done.VAE(os.ReadDir(cfg.SourceCodeRootPath)).Done() {
		if subInfo.IsDir() {
			continue
		}
		if !strings.HasSuffix(subInfo.Name(), ".go") {
			continue
		}
		path := filepath.Join(cfg.SourceCodeRootPath, subInfo.Name())
		zaplog.LOG.Debug(path)

		source := done.VAE(os.ReadFile(path)).Done()

		astFile := done.VCE(utils_golang_ast.NewAstXSourceCode(path, source)).Nice()
		astFns := utils_golang_ast.GetFunctions(astFile)
		mebFunctions := utils_golang_ast.GetFuncsXRecvName(astFns, objectType.Name(), true)
		if len(mebFunctions) == 0 {
			continue
		}

		astTuples = append(astTuples, &astSrcTmpTuple{
			source:       source,
			mebFunctions: mebFunctions,
		})
	}

	if cfg.OptRecvName == "" {
		cfg.OptRecvName = utils.SOrX(astTuples.GetRecvName(), "T")
	}

	if subClassName == "" {
		subClassName = makeNewTypeName(reflect.TypeOf(object), cfg, flexibleType)
	}

	ptx := utils.NewPTX()
	ptx.Println(`type ` + subClassName + ` struct{` + cfg.OptRecvName + ` *` + objectType.Name() + `}`)
	ptx.Println(`
		func(` + cfg.OptRecvName + ` *` + objectType.Name() + `) ` + string(flexibleType) + `() * ` + subClassName + `{
		return & ` + subClassName + `{` + cfg.OptRecvName + `:` + cfg.OptRecvName + `}
	}`)

	for _, oneTmp := range astTuples {
		source := oneTmp.source
		mebFunctions := oneTmp.mebFunctions
		for _, mebFunc := range mebFunctions {
			mebFuncName := utils_golang_ast.GetNodeString(source, mebFunc.Name)
			if utils.In(mebFuncName, []string{string(mustdone.MUST), string(mustdone.SOFT)}) {
				continue
			}

			var params = make(utils_golang_ast_fields.NameTypeElements, 0)
			if mebFunc.Type != nil && mebFunc.Type.Params != nil {
				params = utils_golang_ast_fields.GetSimpleArgElements(mebFunc.Type.Params.List, source)
			}
			var results = make(utils_golang_ast_fields.NameTypeElements, 0)
			if mebFunc.Type != nil && mebFunc.Type.Results != nil {
				results = utils_golang_ast_fields.GetSimpleResElements(mebFunc.Type.Results.List, source)
			}

			for _, elem := range results {
				zaplog.LOG.Debug("elem", zap.String("name", elem.Name), zap.String("kind", elem.Kind))
			}

			var okxResElems = make(utils_golang_ast_fields.NameTypeElements, 0)
			var erxResElems = make(utils_golang_ast_fields.NameTypeElements, 0)
			for _, result := range results {
				if utils_golang_ast.GetNodeString(source, result.Type) == "error" {
					erxResElems = append(erxResElems, result)
				} else {
					okxResElems = append(okxResElems, result)
				}
			}

			var erxHandleStmts []string
			for _, erxElemName := range erxResElems.GetFuncParamsStats() {
				className := cfg.FlexClass
				if className == "" {
					className = mustdone.GetPkgName()
				}
				erxHandleStmts = append(erxHandleStmts, className+"."+string(flexibleType)+"("+erxElemName+")")
			}

			ptx.Println(`func (T *` + subClassName + `) ` + mebFuncName + `(` +
				params.GetNamesKindsStats().Merge() +
				`)` + `(` +
				okxResElems.GetNamesKindsStats().Merge() +
				`) {`)

			runFuncLine := `T.` + cfg.OptRecvName + `.` + mebFuncName + `(` + params.GetFuncParamsStats().Merge() + `)`
			if len(results) > 0 {
				if len(okxResElems) == len(results) {
					ptx.Println(results.GetFuncParamsStats().Merge() + "=" + runFuncLine)
				} else {
					ptx.Println(results.GetFuncParamsStats().Merge() + ":=" + runFuncLine)
				}
			} else {
				ptx.Println(runFuncLine)
			}

			if len(erxHandleStmts) > 0 {
				ptx.Println(strings.Join(erxHandleStmts, "\n"))
			}

			if len(okxResElems) > 0 {
				ptx.Println("return" + " " + okxResElems.GetFuncParamsStats().Merge())
			}

			ptx.Println("}")
		}
	}
	res := ptx.String()
	zaplog.LOG.Debug(res)
	return res
}

type StyleType string

//goland:noinspection GoSnakeCaseUsage
const (
	STYLE_PREFIX_LOWER_TYPE     StyleType = "STYLE_PREFIX_LOWER_TYPE"
	STYLE_PREFIX_UPPER_TYPE     StyleType = "STYLE_PREFIX_UPPER_TYPE"
	STYLE_PREFIX_CAMELCASE_TYPE StyleType = "STYLE_PREFIX_CAMELCASE_TYPE"
	STYLE_SUFFIX_LOWER_TYPE     StyleType = "STYLE_SUFFIX_LOWER_TYPE"
	STYLE_SUFFIX_UPPER_TYPE     StyleType = "STYLE_SUFFIX_UPPER_TYPE"
	STYLE_SUFFIX_CAMELCASE_TYPE StyleType = "STYLE_SUFFIX_CAMELCASE_TYPE"
)

func makeNewTypeName(objectType reflect.Type, cfg *GenParam, flexibleType mustdone.FlexibleHandlingType) string {
	switch cfg.SubClassNameStyleType {
	case STYLE_PREFIX_LOWER_TYPE:
		return strings.ToLower(string(flexibleType)) + cfg.SubClassNamePartWords + objectType.Name()
	case STYLE_PREFIX_UPPER_TYPE:
		return strings.ToUpper(string(flexibleType)) + cfg.SubClassNamePartWords + objectType.Name()
	case STYLE_PREFIX_CAMELCASE_TYPE:
		return string(flexibleType) + cfg.SubClassNamePartWords + objectType.Name()
	case STYLE_SUFFIX_LOWER_TYPE:
		return objectType.Name() + cfg.SubClassNamePartWords + strings.ToLower(string(flexibleType))
	case STYLE_SUFFIX_UPPER_TYPE:
		return objectType.Name() + cfg.SubClassNamePartWords + strings.ToUpper(string(flexibleType))
	case STYLE_SUFFIX_CAMELCASE_TYPE:
		return objectType.Name() + cfg.SubClassNamePartWords + string(flexibleType)
	}
	return strings.ToLower(string(flexibleType)) + cfg.SubClassNamePartWords + objectType.Name()
}

type astSrcTmpTuple struct {
	mebFunctions []*ast.FuncDecl
	source       []byte
}

type astSrcTmpTuples []*astSrcTmpTuple

func (astTuples astSrcTmpTuples) GetRecvName() string {
	for _, oneTmp := range astTuples {
		for _, mebFunction := range oneTmp.mebFunctions {
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
