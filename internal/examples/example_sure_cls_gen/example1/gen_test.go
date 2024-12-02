package example1

import (
	"testing"

	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/sure"
	"github.com/yyle88/sure/sure_cls_gen"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func TestGen(t *testing.T) {
	cfg := &sure_cls_gen.ClassGenConfig{
		ClassGenOptions: sure_cls_gen.NewClassGenOptions(runpath.PARENT.Path()).
			WithNewClassNameParts("88").
			WithNamingPatternType(sure_cls_gen.STYLE_SUFFIX_CAMELCASE_TYPE).
			MoreErrorHandlingModes(sure.MUST, sure.SOFT, sure.OMIT),
		PackageName:   syntaxgo.CurrentPackageName(),
		ImportOptions: syntaxgo_ast.NewPackageImportOptions(),
		OutputPath:    runtestpath.SrcPath(t),
	}

	sure_cls_gen.GenerateClasses(cfg, Example{})
}
