package example6

import (
	"testing"

	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/sure/cls_stub_gen"
	"github.com/yyle88/sure/internal/examples/example6/example6aaa"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func TestGen(t *testing.T) {
	param := cls_stub_gen.NewParam(&example6aaa.A{}, "example6aaa.NODE")

	cfg := &cls_stub_gen.Config{
		SrcRoot:       runpath.PARENT.Join("example6aaa"),
		TargetPkgName: syntaxgo.CurrentPackageName(),
		ImportOptions: syntaxgo_ast.NewPackageImportOptions(),
		TargetSrcPath: runtestpath.SrcPath(t),
		CanCreateFile: false,
	}
	cls_stub_gen.Gen(cfg, param)
}
