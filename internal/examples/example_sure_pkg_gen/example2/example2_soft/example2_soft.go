package example2_soft

import (
	"github.com/yyle88/sure"
	"github.com/yyle88/sure/internal/examples/example_sure_pkg_gen/example2"
)

func GetN() int {
	res0, err := example2.GetN()
	sure.Soft(err)
	return res0
}

func GetS() string {
	res0, err := example2.GetS()
	sure.Soft(err)
	return res0
}
