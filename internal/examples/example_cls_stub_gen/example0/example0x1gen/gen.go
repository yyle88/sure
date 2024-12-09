package example0x1gen

import "github.com/yyle88/sure/internal/examples/example_cls_stub_gen/example0"

func Get() *example0.A {
	return example0.STUB1.Get()
}
func Set(arg string) {
	example0.STUB1.Set(arg)
}
func Add(x int, y int) int {
	return example0.STUB1.Add(x, y)
}
func Sub(x int, y int) (int, error) {
	return example0.STUB1.Sub(x, y)
}
func Who(param ...example0.Param) {
	example0.STUB1.Who(param...)
}
func How(param ...example0.Param) {
	example0.STUB1.How(param...)
}
