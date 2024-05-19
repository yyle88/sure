package example4

import (
	"testing"

	"github.com/yyle88/mustdone/mustsoft_gen_cls"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/syntaxgo"
)

func TestGenCode(t *testing.T) {
	cfg := &mustsoft_gen_cls.Config{
		GenParam:      mustsoft_gen_cls.NewGenParam(runpath.PARENT.Path()),
		PkgName:       syntaxgo.CurrentPackageName(),
		ImportOptions: nil,
		SrcPath:       runtestpath.SrcPath(t),
	}
	mustsoft_gen_cls.Gen(cfg, Example{}, Demo{})
}
