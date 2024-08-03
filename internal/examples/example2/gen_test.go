package example2

import (
	"testing"

	"github.com/yyle88/runpath"
	"github.com/yyle88/sure"
	"github.com/yyle88/sure/mustsoft_gen_pkg"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

type ObjectType struct{}

func TestGen(t *testing.T) {
	pkgPath := syntaxgo_reflect.GetPkgPathV2[ObjectType]()
	mustsoft_gen_pkg.Gen(t, runpath.PARENT.Path(), sure.SOFT, pkgPath)
	mustsoft_gen_pkg.Gen(t, runpath.PARENT.Path(), sure.MUST, pkgPath)
}
