package example3_must

import (
	"github.com/yyle88/sure"
	"github.com/yyle88/sure/internal/examples/example3"
)

func Neat[T any](a T) []byte {
	res0, err := example3.Neat[T](a)
	sure.Must(err)
	return res0
}

func Bind[T any](data []byte) *T {
	res0, err := example3.Bind[T](data)
	sure.Must(err)
	return res0
}
