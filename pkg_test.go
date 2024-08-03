package sure_test

import (
	"testing"

	"github.com/yyle88/sure"
)

func TestGetPkgPath(t *testing.T) {
	t.Log(sure.GetPkgPath())
}

func TestGetPkgName(t *testing.T) {
	t.Log(sure.GetPkgName())
}
