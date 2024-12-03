package example6

type A struct {
	name string
}

func NewA(name string) *A {
	return &A{name: name}
}

func (a *A) Name() string {
	return a.name
}
