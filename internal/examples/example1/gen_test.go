package example1

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/formatgo"
	"github.com/yyle88/mustdone"
	"github.com/yyle88/mustdone/internal/utils"
	"github.com/yyle88/mustdone/mustsoft_gen_cls"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

func TestGenCode(t *testing.T) {
	root := runpath.PARENT.Path()
	packageName := syntaxgo.CurrentPackageName()

	ptx := utils.NewPTX()
	ptx.Println("package", packageName)

	param := &mustsoft_gen_cls.GenParam{
		SourceCodeRootPath:    root,
		SubClassNamePartWords: "88",
		SubClassNameStyleType: mustsoft_gen_cls.STYLE_SUFFIX_CAMELCASE_TYPE,
	}
	ptx.Println(mustsoft_gen_cls.GenCodes(Example{}, param, mustdone.MUST, mustdone.SOFT))

	//很明显，当主动写 import 的时候，执行的 format 的速度特别快
	packageImportOptions := &syntaxgo_ast.PackageImportOptions{
		Packages:   nil,
		UsingTypes: nil,
		Objects: []any{
			syntaxgo_reflect.GetObject[mustdone.FlexibleEnum](),
		},
	}
	source := syntaxgo_ast.AddImports(ptx.Bytes(), packageImportOptions)
	//假如你不在写文件的时候 format，代码格式就不美观，而执行 format 时，最后确保它不再去找包名，因为当项目过大时找包名会变得很困难的
	newSource, err := formatgo.FormatBytes(source)
	require.NoError(t, err)

	srcPath := runtestpath.SrcPath(t)
	require.NoError(t, utils.WriteFile(srcPath, newSource))
}
