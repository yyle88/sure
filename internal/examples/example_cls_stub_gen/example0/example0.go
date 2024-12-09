package example0

var STUB1 = &A{}
var STUB2 = &A{}

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

type Param struct{}

func (a *A) Who(param ...Param) {}

func (a *A) How(param ...Param) {}
