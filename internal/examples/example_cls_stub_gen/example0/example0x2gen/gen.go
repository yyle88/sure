package example0x2gen

import "github.com/yyle88/sure/internal/examples/example_cls_stub_gen/example0"

func Get() *example0.A {
	return example0.STUB2.Get()
}
func Set(arg string) {
	example0.STUB2.Set(arg)
}
func Add(x int, y int) int {
	return example0.STUB2.Add(x, y)
}
func Sub(x int, y int) (int, error) {
	return example0.STUB2.Sub(x, y)
}
