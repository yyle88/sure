package example5

import (
	"testing"

	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/sure/internal/examples/example5/example5mustsoftnode"
	"github.com/yyle88/sure/mustsoft_gen_cls"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func TestGen(t *testing.T) {
	param := mustsoft_gen_cls.NewGenParam(runpath.PARENT.Path())
	param.SetMustSoftCallableNode("example5mustsoftnode.NODE")

	importOptions := syntaxgo_ast.NewPackageImportOptions()
	importOptions.SetObject(example5mustsoftnode.Node{})

	cfg := &mustsoft_gen_cls.Config{
		GenParam:      param,
		PkgName:       syntaxgo.CurrentPackageName(),
		ImportOptions: importOptions,
		SrcPath:       runtestpath.SrcPath(t),
	}
	mustsoft_gen_cls.Gen(cfg, Example{}, Demo{})
}
