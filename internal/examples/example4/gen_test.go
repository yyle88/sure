package example4

import (
	"testing"

	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/sure/mustsoft_gen_cls"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func TestGen(t *testing.T) {
	param := mustsoft_gen_cls.NewGenParam(runpath.PARENT.Path())
	param.SetMustSoftCallableNode("done")

	importOptions := syntaxgo_ast.NewPackageImportOptions()
	importOptions.SetPkgPath("github.com/yyle88/done")

	cfg := &mustsoft_gen_cls.Config{
		GenParam:      param,
		PkgName:       syntaxgo.CurrentPackageName(),
		ImportOptions: importOptions,
		SrcPath:       runtestpath.SrcPath(t),
	}
	mustsoft_gen_cls.Gen(cfg, Example{}, Demo{})
}
