package example5

import (
	"testing"

	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/sure/internal/examples/example_sure_cls_gen/example5/example5surenode"
	"github.com/yyle88/sure/sure_cls_gen"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func TestGen(t *testing.T) {
	cfg := &sure_cls_gen.ClassGenConfig{
		ClassGenOptions: sure_cls_gen.NewClassGenOptions(runpath.PARENT.Path()).
			WithErrorHandlerFuncName("example5surenode.NODE"),
		PackageName: syntaxgo.CurrentPackageName(),
		ImportOptions: syntaxgo_ast.NewPackageImportOptions().
			SetInferredObject(example5surenode.Node{}),
		OutputPath: runtestpath.SrcPath(t),
	}
	sure_cls_gen.GenerateClasses(cfg, Example{}, Demo{})
}
