package mustdone_test

import (
	"reflect"
	"testing"

	"github.com/yyle88/mustdone"
)

func TestObjectType(t *testing.T) {
	t.Log(reflect.TypeOf(mustdone.ObjectType{}).PkgPath())
}

func TestGetPkgPath(t *testing.T) {
	t.Log(mustdone.GetPkgPath())
}

func TestGetPkgName(t *testing.T) {
	t.Log(mustdone.GetPkgName())
}
