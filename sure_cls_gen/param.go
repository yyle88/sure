package sure_cls_gen

import (
	"reflect"
	"strings"

	"github.com/yyle88/sure"
	"github.com/yyle88/tern"
	"github.com/yyle88/tern/zerotern"
)

type StyleEnum string

//goland:noinspection GoSnakeCaseUsage
const (
	STYLE_PREFIX_LOWER_TYPE StyleEnum = "STYLE_PREFIX_LOWER_TYPE"
	STYLE_SUFFIX_LOWER_TYPE StyleEnum = "STYLE_SUFFIX_LOWER_TYPE"

	STYLE_PREFIX_UPPER_TYPE StyleEnum = "STYLE_PREFIX_UPPER_TYPE"
	STYLE_SUFFIX_UPPER_TYPE StyleEnum = "STYLE_SUFFIX_UPPER_TYPE"

	STYLE_PREFIX_CAMELCASE_TYPE StyleEnum = "STYLE_PREFIX_CAMELCASE_TYPE"
	STYLE_SUFFIX_CAMELCASE_TYPE StyleEnum = "STYLE_SUFFIX_CAMELCASE_TYPE"
)

type GenParam struct {
	SrcRoot               string    //这是必填参数，你要解析的类型所在源代码的目录（目前似乎没有能力通过object就得到代码目录）
	SubClassName          string    //当只需要生成一个类的时候，当然是可以直接设置类名的，否则就要使用下面的字段配置生成的规则
	SubClassNamePartWords string    //非必填参数，你要生成的新子类型的名称片段，就是这个名称字符串中间，有部分可以自定义的内容
	SubClassNameStyleEnum StyleEnum //非必填参数，你要生成的新子类型的命名风格，有默认风格
	SubClassRecvName      string    //默认不填，你要解析的类型它的成员函数的recv的名称，比如 func (a *A)do() 就填写 a 就行
	SureNode              string    //非必填参数，就是调用 SURE 函数的调用者，你也可以实现自己的 sure 函数，默认用 sure 包的
	SureEnums             []sure.SureEnum
}

func NewGenParam(srcRoot string) *GenParam {
	return &GenParam{SrcRoot: srcRoot}
}

func (cfg *GenParam) SetSubClassName(subClassName string) *GenParam {
	cfg.SubClassName = subClassName
	return cfg
}

func (cfg *GenParam) SetSubClassNamePartWords(subClassNamePartWords string) *GenParam {
	cfg.SubClassNamePartWords = subClassNamePartWords
	return cfg
}

func (cfg *GenParam) SetSubClassNameStyleEnum(subClassNameStyleType StyleEnum) *GenParam {
	cfg.SubClassNameStyleEnum = subClassNameStyleType
	return cfg
}

func (cfg *GenParam) SetSubClassRecvName(subClassRecvName string) *GenParam {
	cfg.SubClassRecvName = subClassRecvName
	return cfg
}

func (cfg *GenParam) SetSureNode(sureNode string) *GenParam {
	cfg.SureNode = sureNode
	return cfg
}

func (cfg *GenParam) SetSureEnum(sureEnum ...sure.SureEnum) *GenParam {
	cfg.SureEnums = append(cfg.SureEnums, sureEnum...)
	return cfg
}

func (cfg *GenParam) GetSureEnums() []sure.SureEnum {
	return tern.BVF(len(cfg.SureEnums) != 0, cfg.SureEnums, func() []sure.SureEnum {
		return []sure.SureEnum{sure.MUST, sure.SOFT} //当没有设置时，返回默认的两个最主要的
	})
}

func (cfg *GenParam) makeClassName(objectType reflect.Type, sureEnum sure.SureEnum) string {
	return zerotern.VF(cfg.SubClassName, func() string {
		switch cfg.SubClassNameStyleEnum {
		case STYLE_PREFIX_LOWER_TYPE:
			return strings.ToLower(string(sureEnum)) + cfg.SubClassNamePartWords + objectType.Name()
		case STYLE_SUFFIX_LOWER_TYPE:
			return objectType.Name() + cfg.SubClassNamePartWords + strings.ToLower(string(sureEnum))

		case STYLE_PREFIX_UPPER_TYPE:
			return strings.ToUpper(string(sureEnum)) + cfg.SubClassNamePartWords + objectType.Name()
		case STYLE_SUFFIX_UPPER_TYPE:
			return objectType.Name() + cfg.SubClassNamePartWords + strings.ToUpper(string(sureEnum))

		case STYLE_PREFIX_CAMELCASE_TYPE:
			return string(sureEnum) + cfg.SubClassNamePartWords + objectType.Name()
		case STYLE_SUFFIX_CAMELCASE_TYPE, StyleEnum(""): //默认值就是 ClassNameMust 或者 ClassNameSoft 新类名
			return objectType.Name() + cfg.SubClassNamePartWords + string(sureEnum)
		}
		return strings.ToLower(string(sureEnum)) + cfg.SubClassNamePartWords + objectType.Name()
	})
}
