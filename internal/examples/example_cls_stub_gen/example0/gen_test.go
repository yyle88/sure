package example0

import (
	"testing"

	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/sure/cls_stub_gen"
	"github.com/yyle88/sure/internal/examples/example_cls_stub_gen/example0/example0node"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func TestGen(t *testing.T) {
	param := cls_stub_gen.NewStubParam(&example0node.A{}, "example0node.NODE")

	cfg := &cls_stub_gen.StubGenConfig{
		SourceRootPath:    runpath.PARENT.Join("example0node"),
		TargetPackageName: syntaxgo.CurrentPackageName(),
		ImportOptions:     syntaxgo_ast.NewPackageImportOptions(),
		OutputPath:        runtestpath.SrcPath(t),
		AllowFileCreation: false,
	}
	cls_stub_gen.GenerateStubsFile(cfg, param)
}
