package mustsoft_gen_cls

import (
	"reflect"
	"strings"

	"github.com/yyle88/mustdone"
)

type StyleType string

//goland:noinspection GoSnakeCaseUsage
const (
	STYLE_PREFIX_LOWER_TYPE StyleType = "STYLE_PREFIX_LOWER_TYPE"
	STYLE_SUFFIX_LOWER_TYPE StyleType = "STYLE_SUFFIX_LOWER_TYPE"

	STYLE_PREFIX_UPPER_TYPE StyleType = "STYLE_PREFIX_UPPER_TYPE"
	STYLE_SUFFIX_UPPER_TYPE StyleType = "STYLE_SUFFIX_UPPER_TYPE"

	STYLE_PREFIX_CAMELCASE_TYPE StyleType = "STYLE_PREFIX_CAMELCASE_TYPE"
	STYLE_SUFFIX_CAMELCASE_TYPE StyleType = "STYLE_SUFFIX_CAMELCASE_TYPE"
)

type GenParam struct {
	SrcRoot               string    //这是必填参数，你要解析的类型所在源代码的目录（目前似乎没有能力通过object就得到代码目录）
	SubClassName          string    //当只需要生成一个类的时候，当然是可以直接设置类名的，否则就要使用下面的字段配置生成的规则
	SubClassNamePartWords string    //非必填参数，你要生成的新子类型的名称片段，就是这个名称字符串中间，有部分可以自定义的内容
	SubClassNameStyleType StyleType //非必填参数，你要生成的新子类型的命名风格，有默认风格
	OptRecvName           string    //默认不填，你要解析的类型它的成员函数的recv的名称，比如 func (a *A)do() 就填写 a 就行
	FlexibleClass         string    //非必填参数，就是调用 FLEX 函数的调用者，你也可以实现自己的 flex 函数，默认用 flex 包的
}

func NewGenParam(srcRoot string) *GenParam {
	return &GenParam{SrcRoot: srcRoot}
}

func (cfg *GenParam) makeClassName(objectType reflect.Type, flexibleEnum mustdone.FlexibleEnum) string {
	if cfg.SubClassName != "" {
		return cfg.SubClassName
	}

	switch cfg.SubClassNameStyleType {
	case STYLE_PREFIX_LOWER_TYPE:
		return strings.ToLower(string(flexibleEnum)) + cfg.SubClassNamePartWords + objectType.Name()
	case STYLE_SUFFIX_LOWER_TYPE:
		return objectType.Name() + cfg.SubClassNamePartWords + strings.ToLower(string(flexibleEnum))

	case STYLE_PREFIX_UPPER_TYPE:
		return strings.ToUpper(string(flexibleEnum)) + cfg.SubClassNamePartWords + objectType.Name()
	case STYLE_SUFFIX_UPPER_TYPE:
		return objectType.Name() + cfg.SubClassNamePartWords + strings.ToUpper(string(flexibleEnum))

	case STYLE_PREFIX_CAMELCASE_TYPE:
		return string(flexibleEnum) + cfg.SubClassNamePartWords + objectType.Name()
	case STYLE_SUFFIX_CAMELCASE_TYPE, StyleType(""): //默认值就是 ClassNameMust 或者 ClassNameSoft 新类名
		return objectType.Name() + cfg.SubClassNamePartWords + string(flexibleEnum)
	}
	return strings.ToLower(string(flexibleEnum)) + cfg.SubClassNamePartWords + objectType.Name()
}