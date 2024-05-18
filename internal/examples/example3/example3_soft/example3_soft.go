package example3_soft

import (
	"github.com/yyle88/mustdone"
	"github.com/yyle88/mustdone/internal/examples/example3"
)

func Neat[T any](a T) []byte {
	res0, err := example3.Neat[T](a)
	mustdone.Soft(err)
	return res0
}

func Bind[T any](data []byte) *T {
	res0, err := example3.Bind[T](data)
	mustdone.Soft(err)
	return res0
}
