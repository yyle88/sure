package examples

import (
	"math/rand/v2"

	"github.com/yyle88/sure/internal/examples/example_cls_stub_gen/example0"
)

func Example0() int {
	return example0.Add(rand.IntN(100), rand.IntN(100))
}
