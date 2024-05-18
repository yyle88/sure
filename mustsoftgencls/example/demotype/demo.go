package demotype

type Demo struct {
	value int
}

func NewDemo(value int) *Demo {
	return &Demo{value: value}
}

func (a *Demo) GetInt() (int, error) {
	return int(a.value), nil
}

func (a *Demo) GetFloat64() (float64, error) {
	return float64(a.value), nil
}
