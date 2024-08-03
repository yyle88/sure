package example2_must

import (
	"github.com/yyle88/sure"
	"github.com/yyle88/sure/internal/examples/example2"
)

func GetN() int {
	res0, err := example2.GetN()
	sure.Must(err)
	return res0
}

func GetS() string {
	res0, err := example2.GetS()
	sure.Must(err)
	return res0
}
