package example0

import "github.com/yyle88/sure/internal/examples/example_cls_stub_gen/example0/example0node"

func Get() *example0node.A {
	return example0node.NODE.Get()
}
func Set(arg string) {
	example0node.NODE.Set(arg)
}
func Add(x int, y int) int {
	return example0node.NODE.Add(x, y)
}
func Sub(x int, y int) (int, error) {
	return example0node.NODE.Sub(x, y)
}
