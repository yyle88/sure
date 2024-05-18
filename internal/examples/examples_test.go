package examples

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/mustdone/internal/examples/example1"
	"github.com/yyle88/mustdone/internal/examples/example2/example2_must"
	"github.com/yyle88/mustdone/internal/examples/example2/example2_soft"
	"github.com/yyle88/mustdone/internal/examples/example3/example3_must"
	"github.com/yyle88/mustdone/internal/examples/example3/example3_soft"
)

func TestExample1(t *testing.T) {
	//当你要操作的是个对象时，你就可以赋予这个对象 Must 和 Soft 的能力，Must表示出错时崩溃，而Soft表示出错是仅仅告警但流程继续
	a := example1.NewExample(1, "s")
	t.Log(a.Must().GetN())
	t.Log(a.Must().GetS())

	t.Log(a.Soft().GetN())
	t.Log(a.Soft().GetS())
}

func TestExample2(t *testing.T) {
	//但是假如你要操作的是个包，你就可以在这个包的基础上衍生出 must 包 和 soft 包，能让你在调用时避免判断错误
	t.Log(example2_must.GetN())
	t.Log(example2_must.GetS())

	t.Log(example2_soft.GetN())
	t.Log(example2_soft.GetS())
}

func TestExample3(t *testing.T) {
	//在泛型的情况下依然是可以的，我想这正是我们需要的，特别是下面这俩函数，把对象转化为json和把json转化为对象
	type Param struct {
		Name string
	}

	{ //使用 soft 软包把对象转json，再用 must 硬包把json转为对象
		data := example3_soft.Neat(&Param{Name: "haha"})
		t.Log(string(data))

		resX := example3_must.Bind[Param](data)
		require.NotNil(t, resX) //其实这块是不需要判断的，毕竟是must的，保证是成功的
		require.Equal(t, "haha", resX.Name)
	}

	{ //当然也可以反过来用，也可以都用 soft 软的也可以都用 must 硬的，这里只是演示，根据场景自由选择
		data := example3_must.Neat(&Param{Name: "haha"})
		t.Log(string(data))

		resX := example3_soft.Bind[Param](data)
		require.NotNil(t, resX) //这块是需要判读的，毕竟是soft的，仅仅是忽略错误发出告警但流程继续执行
		require.Equal(t, "haha", resX.Name)
	}
}
