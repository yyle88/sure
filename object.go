package mustdone

import (
	"reflect"

	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

type ObjectType struct{}

func GetPkgPath() string {
	return reflect.TypeOf(ObjectType{}).PkgPath()
}

func GetPkgName() string {
	return syntaxgo_reflect.GetPkgName(ObjectType{})
}
