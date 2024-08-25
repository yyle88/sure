package example6

import "github.com/yyle88/sure/internal/examples/example6/example6aaa"

func Get() *example6aaa.A {
	return example6aaa.NODE.Get()
}
func Set(arg string) {
	example6aaa.NODE.Set(arg)
}
func Add(x int, y int) int {
	return example6aaa.NODE.Add(x, y)
}
func Sub(x int, y int) (int, error) {
	return example6aaa.NODE.Sub(x, y)
}
