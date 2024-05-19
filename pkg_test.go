package mustdone_test

import (
	"testing"

	"github.com/yyle88/mustdone"
)

func TestGetPkgPath(t *testing.T) {
	t.Log(mustdone.GetPkgPath())
}

func TestGetPkgName(t *testing.T) {
	t.Log(mustdone.GetPkgName())
}
