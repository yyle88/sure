package example3_must

import (
	"github.com/yyle88/mustdone"
	"github.com/yyle88/mustdone/internal/examples/example3"
)

func Neat[T any](a T) []byte {
	res0, err := example3.Neat[T](a)
	mustdone.Must(err)
	return res0
}

func Bind[T any](data []byte) *T {
	res0, err := example3.Bind[T](data)
	mustdone.Must(err)
	return res0
}
