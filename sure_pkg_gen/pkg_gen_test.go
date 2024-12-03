package sure_pkg_gen

import (
	"testing"

	"github.com/yyle88/runpath"
	"github.com/yyle88/sure"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

func TestMain(m *testing.M) {
	m.Run()
}

type ObjectType struct{}

func DemoFunction(objectType *ObjectType) *ObjectType {
	return objectType
}

func TestGenerateSureFunctions(t *testing.T) {
	path := runpath.Path()

	pkgPath := syntaxgo_reflect.GetPkgPathV2[ObjectType]()

	config := NewSurePackageConfig(runpath.PARENT.Path(), sure.MUST, pkgPath)

	sureFunctions := GenerateSureFunctions(t, config, path)

	for _, sureFunction := range sureFunctions {
		t.Log(sureFunction)
	}
}
