package example6x2gen

import (
	"testing"

	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/sure/cls_stub_gen"
	"github.com/yyle88/sure/internal/examples/example_cls_stub_gen/example6"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func TestGen(t *testing.T) {
	param := cls_stub_gen.NewStubParam(&example6.A{}, "A200")

	cfg := &cls_stub_gen.StubGenConfig{
		SourceRootPath:    runpath.PARENT.Join("../../example6"),
		TargetPackageName: syntaxgo.CurrentPackageName(),
		ImportOptions:     syntaxgo_ast.NewPackageImportOptions(),
		OutputPath:        runtestpath.SrcPath(t),
		AllowFileCreation: false,
	}
	cls_stub_gen.GenerateStubs(cfg, param)
}
