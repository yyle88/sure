package example2

import (
	"testing"

	"github.com/yyle88/runpath"
	"github.com/yyle88/sure"
	"github.com/yyle88/sure/sure_pkg_gen"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

type ObjectType struct{}

func TestGen(t *testing.T) {
	pkgPath := syntaxgo_reflect.GetPkgPathV2[ObjectType]()
	sure_pkg_gen.Gen(t, runpath.PARENT.Path(), sure.SOFT, pkgPath)
	sure_pkg_gen.Gen(t, runpath.PARENT.Path(), sure.MUST, pkgPath)
	sure_pkg_gen.Gen(t, runpath.PARENT.Path(), sure.OMIT, pkgPath)
}
