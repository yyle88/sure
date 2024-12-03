package example_cls_stub_gen

import (
	"math/rand/v2"

	"github.com/yyle88/sure/internal/examples/example_cls_stub_gen/example0/example0x1gen"
	"github.com/yyle88/sure/internal/examples/example_cls_stub_gen/example0/example0x2gen"
	"github.com/yyle88/sure/internal/examples/example_cls_stub_gen/example6/example6x1gen"
	"github.com/yyle88/sure/internal/examples/example_cls_stub_gen/example6/example6x2gen"
)

func Example0x1() int {
	return example0x1gen.Add(rand.IntN(100), rand.IntN(100))
}

func Example0x2() int {
	return example0x2gen.Add(rand.IntN(100), rand.IntN(100))
}

func Example6x1() string {
	return example6x1gen.Name()
}

func Example6x2() string {
	return example6x2gen.Name()
}
