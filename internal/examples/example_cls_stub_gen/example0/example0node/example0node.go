package example0node

var NODE = &A{}

type A struct{}

func (a *A) Get() *A {
	return a
}

func (a *A) Set(string) {
}

func (a *A) Add(x int, y int) int {
	return x + y
}

func (a *A) Sub(x int, y int) (int, error) {
	return x - y, nil
}
