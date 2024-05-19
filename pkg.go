package mustdone

import (
	"reflect"

	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

func GetPkgPath() string {
	return reflect.TypeOf(syntaxgo_reflect.GetObject[FlexibleEnum]()).PkgPath()
}

func GetPkgName() string {
	return syntaxgo_reflect.GetPkgName(syntaxgo_reflect.GetObject[FlexibleEnum]())
}
