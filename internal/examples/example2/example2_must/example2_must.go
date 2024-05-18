package example2_must

import (
	"github.com/yyle88/mustdone"
	"github.com/yyle88/mustdone/internal/examples/example2"
)

func GetN() int {
	res0, err := example2.GetN()
	mustdone.Must(err)
	return res0
}

func GetS() string {
	res0, err := example2.GetS()
	mustdone.Must(err)
	return res0
}
