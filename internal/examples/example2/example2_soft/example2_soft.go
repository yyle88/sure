package example2_soft

import (
	"github.com/yyle88/mustdone"
	"github.com/yyle88/mustdone/internal/examples/example2"
)

func GetN() int {
	res0, err := example2.GetN()
	mustdone.Soft(err)
	return res0
}

func GetS() string {
	res0, err := example2.GetS()
	mustdone.Soft(err)
	return res0
}
