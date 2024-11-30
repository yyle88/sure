package example5

import (
	"testing"

	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/sure/internal/examples/example5/example5surenode"
	"github.com/yyle88/sure/sure_cls_gen"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func TestGen(t *testing.T) {
	param := sure_cls_gen.NewGenParam(runpath.PARENT.Path())
	param.SetSureNode("example5surenode.NODE")

	importOptions := syntaxgo_ast.NewPackageImportOptions()
	importOptions.SetInferredObject(example5surenode.Node{})

	cfg := &sure_cls_gen.Config{
		GenParam:      param,
		PkgName:       syntaxgo.CurrentPackageName(),
		ImportOptions: importOptions,
		SrcPath:       runtestpath.SrcPath(t),
	}
	sure_cls_gen.Gen(cfg, Example{}, Demo{})
}
