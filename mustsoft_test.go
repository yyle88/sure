package mustdone

import (
	"testing"

	"github.com/pkg/errors"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestMust(t *testing.T) {
	t.Log("-")
	Must(nil) //当没有错误时就什么也不做，当非空时将会崩溃
	t.Log("-")
}

func TestSoft(t *testing.T) {
	t.Log("-")
	Soft(errors.New("wrong")) //将会告警，而且程序将继续执行
	t.Log("-")
}
