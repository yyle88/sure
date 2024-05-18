package example3

import (
	"testing"

	"github.com/yyle88/mustdone"
	"github.com/yyle88/mustdone/mustsoft_gen_pkg"
	"github.com/yyle88/runpath"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

type ObjectType struct{}

func TestGen(t *testing.T) {
	pkgPath := syntaxgo_reflect.GetPkgPathV2[ObjectType]()

	//选择生成 soft 软包
	mustsoft_gen_pkg.Gen(t, runpath.PARENT.Path(), mustdone.SOFT, pkgPath)
	//选择生成 must 硬包
	mustsoft_gen_pkg.Gen(t, runpath.PARENT.Path(), mustdone.MUST, pkgPath)
}
