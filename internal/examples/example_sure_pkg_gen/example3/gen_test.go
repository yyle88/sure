package example3

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

	//选择生成 soft 软包
	sure_pkg_gen.GenerateSurePackage(t, runpath.PARENT.Path(), sure.SOFT, pkgPath)
	//选择生成 must 硬包
	sure_pkg_gen.GenerateSurePackage(t, runpath.PARENT.Path(), sure.MUST, pkgPath)
}
