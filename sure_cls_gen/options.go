package sure_cls_gen

import (
	"reflect"
	"strings"

	"github.com/yyle88/sure"
	"github.com/yyle88/tern"
	"github.com/yyle88/tern/zerotern"
)

type NamingPattern string

//goland:noinspection GoSnakeCaseUsage
const (
	STYLE_PREFIX_LOWER_TYPE NamingPattern = "STYLE_PREFIX_LOWER_TYPE"
	STYLE_SUFFIX_LOWER_TYPE NamingPattern = "STYLE_SUFFIX_LOWER_TYPE"

	STYLE_PREFIX_UPPER_TYPE NamingPattern = "STYLE_PREFIX_UPPER_TYPE"
	STYLE_SUFFIX_UPPER_TYPE NamingPattern = "STYLE_SUFFIX_UPPER_TYPE"

	STYLE_PREFIX_CAMELCASE_TYPE NamingPattern = "STYLE_PREFIX_CAMELCASE_TYPE"
	STYLE_SUFFIX_CAMELCASE_TYPE NamingPattern = "STYLE_SUFFIX_CAMELCASE_TYPE"
)

type ClassGenOptions struct {
	SourceRootPath       string        //这是必填参数，你要解析的类型所在源代码的目录（目前似乎没有能力通过object就得到代码目录）
	NewClassName         string        //当只需要生成一个类的时候，当然是可以直接设置类名的，否则就要使用下面的字段配置生成的规则
	NewClassNameParts    string        //非必填参数，你要生成的新子类型的名称片段，就是这个名称字符串中间，有部分可以自定义的内容
	NamingPatternType    NamingPattern //非必填参数，你要生成的新子类型的命名风格，有默认风格
	ReceiverVariableName string        //默认不填，你要解析的类型它的成员函数的recv的名称，比如 func (a *A)do() 就填写 a 就行
	ErrorHandlerFuncName string        //非必填参数，就是调用 SURE 函数的调用者，你也可以实现自己的 sure 函数，默认用 sure 包的
	ErrorHandlingModes   []sure.ErrorHandlingMode
}

func NewClassGenOptions(srcRoot string) *ClassGenOptions {
	return &ClassGenOptions{SourceRootPath: srcRoot}
}

func (cfg *ClassGenOptions) WithNewClassName(newClassName string) *ClassGenOptions {
	cfg.NewClassName = newClassName
	return cfg
}

func (cfg *ClassGenOptions) WithNewClassNameParts(newClassNameParts string) *ClassGenOptions {
	cfg.NewClassNameParts = newClassNameParts
	return cfg
}

func (cfg *ClassGenOptions) WithNamingPatternType(namingPatternType NamingPattern) *ClassGenOptions {
	cfg.NamingPatternType = namingPatternType
	return cfg
}

func (cfg *ClassGenOptions) WithReceiverVariableName(receiverVariableName string) *ClassGenOptions {
	cfg.ReceiverVariableName = receiverVariableName
	return cfg
}

func (cfg *ClassGenOptions) WithErrorHandlerFuncName(errorHandlerFuncName string) *ClassGenOptions {
	cfg.ErrorHandlerFuncName = errorHandlerFuncName
	return cfg
}

func (cfg *ClassGenOptions) MoreErrorHandlingModes(errorHandlingModes ...sure.ErrorHandlingMode) *ClassGenOptions {
	cfg.ErrorHandlingModes = append(cfg.ErrorHandlingModes, errorHandlingModes...)
	return cfg
}

func (cfg *ClassGenOptions) GetErrorHandlingModes() []sure.ErrorHandlingMode {
	return tern.BVF(len(cfg.ErrorHandlingModes) != 0, cfg.ErrorHandlingModes, func() []sure.ErrorHandlingMode {
		return []sure.ErrorHandlingMode{sure.MUST, sure.SOFT} //当没有设置时，返回默认的两个最主要的
	})
}

func (cfg *ClassGenOptions) GenerateNewClassName(objectType reflect.Type, sureFlag sure.ErrorHandlingMode) string {
	return zerotern.VF(cfg.NewClassName, func() string {
		switch cfg.NamingPatternType {
		case STYLE_PREFIX_LOWER_TYPE:
			return strings.ToLower(string(sureFlag)) + cfg.NewClassNameParts + objectType.Name()
		case STYLE_SUFFIX_LOWER_TYPE:
			return objectType.Name() + cfg.NewClassNameParts + strings.ToLower(string(sureFlag))

		case STYLE_PREFIX_UPPER_TYPE:
			return strings.ToUpper(string(sureFlag)) + cfg.NewClassNameParts + objectType.Name()
		case STYLE_SUFFIX_UPPER_TYPE:
			return objectType.Name() + cfg.NewClassNameParts + strings.ToUpper(string(sureFlag))

		case STYLE_PREFIX_CAMELCASE_TYPE:
			return string(sureFlag) + cfg.NewClassNameParts + objectType.Name()
		case STYLE_SUFFIX_CAMELCASE_TYPE, NamingPattern(""): //默认值就是 ClassNameMust 或者 ClassNameSoft 新类名
			return objectType.Name() + cfg.NewClassNameParts + string(sureFlag)
		}
		return strings.ToLower(string(sureFlag)) + cfg.NewClassNameParts + objectType.Name()
	})
}
