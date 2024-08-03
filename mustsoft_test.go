package sure_test

import (
	"encoding/json"
	"testing"

	"github.com/pkg/errors"
	"github.com/yyle88/sure"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestMust(t *testing.T) {
	type Example struct {
		Name string `json:"name"`
	}

	example := &Example{Name: "lele"}
	t.Log(example)

	data, err := json.Marshal(example)
	sure.Must(err) //当没有错误时就什么也不做，当出错时将 panic 崩溃

	t.Log(string(data))
}

func TestSoft(t *testing.T) {
	t.Log("-")
	sure.Soft(errors.New("wrong")) //将会告警，而且程序将继续执行
	t.Log("-")
}
