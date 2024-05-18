package demotype

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/formatgo"
	"github.com/yyle88/mustdone"
	"github.com/yyle88/mustdone/internal/utils"
	"github.com/yyle88/mustdone/mustsoftgencls"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func TestGenCode(t *testing.T) {
	root := runpath.PARENT.Path()
	packageName := runpath.PARENT.Name()

	ptx := utils.NewPTX()
	ptx.Println("package", packageName)
	param := &mustsoftgencls.GenParam{
		SourceCodeRootPath:    root,
		SubClassNamePartWords: "88",
		SubClassNameStyleType: mustsoftgencls.STYLE_SUFFIX_CAMELCASE_TYPE,
	}
	ptx.Println(mustsoftgencls.GenCodes(Demo{}, param, mustdone.MUST, mustdone.SOFT))

	srcPath := runtestpath.SrcPath(t)
	source := ptx.Bytes()
	//得到空的对象
	var object1 mustdone.FlexibleHandlingType
	var object2 os.File
	//很明显，当主动写 import 的时候，执行的 format 的速度特别快
	source = syntaxgo_ast.AddImportsOfObjects(srcPath, source, []any{object1, object2})
	//假如你不在写文件的时候 format，代码格式就不美观，而执行 format 时，最后确保它不再去找包名，因为当项目过大时找包名会变得很困难的
	newSource, err := formatgo.FormatBytes(source)
	require.NoError(t, err)
	require.NoError(t, utils.WriteFile(srcPath, newSource))
}
