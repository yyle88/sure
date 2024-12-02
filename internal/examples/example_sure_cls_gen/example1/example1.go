package example1

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
