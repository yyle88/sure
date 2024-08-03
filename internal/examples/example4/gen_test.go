package example4

import (
	"testing"

	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/sure/sure_cls_gen"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func TestGen(t *testing.T) {
	param := sure_cls_gen.NewGenParam(runpath.PARENT.Path())
	param.SetSureNode("done")

	importOptions := syntaxgo_ast.NewPackageImportOptions()
	importOptions.SetPkgPath("github.com/yyle88/done")

	cfg := &sure_cls_gen.Config{
		GenParam:      param,
		PkgName:       syntaxgo.CurrentPackageName(),
		ImportOptions: importOptions,
		SrcPath:       runtestpath.SrcPath(t),
	}
	sure_cls_gen.Gen(cfg, Example{}, Demo{})
}
