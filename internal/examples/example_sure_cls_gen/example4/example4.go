package example4

type Example struct {
	n int
	s string
}

func NewExample(n int, s string) *Example {
	return &Example{
		n: n,
		s: s,
	}
}

func (a *Example) GetN() (int, error) {
	return a.n, nil
}

func (a *Example) GetS() (string, error) {
	return a.s, nil
}

type Demo struct {
	n int
	s string
}

func NewDemo(n int, s string) *Demo {
	return &Demo{
		n: n,
		s: s,
	}
}

func (a *Demo) GetN() (int, error) {
	return a.n, nil
}

func (a *Demo) GetS() (string, error) {
	return a.s, nil
}
