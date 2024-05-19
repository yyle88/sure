package example4

import (
	"testing"

	"github.com/yyle88/mustdone/mustsoft_gen_cls"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func TestGenCode(t *testing.T) {
	cfg := &mustsoft_gen_cls.Config{
		GenParam:      mustsoft_gen_cls.NewGenParam(runpath.PARENT.Path()),
		PkgName:       syntaxgo.CurrentPackageName(),
		ImportOptions: syntaxgo_ast.NewPackageImportOptions(),
		SrcPath:       runtestpath.SrcPath(t),
	}
	mustsoft_gen_cls.Gen(cfg, Example{}, Demo{})
}
