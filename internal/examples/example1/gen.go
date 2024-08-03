package example1

import "github.com/yyle88/sure"

type Example88Must struct{ a *Example }

func (a *Example) Must() *Example88Must {
	return &Example88Must{a: a}
}
func (T *Example88Must) GetN() (res int) {
	res, err1 := T.a.GetN()
	sure.Must(err1)
	return res
}
func (T *Example88Must) GetS() (res string) {
	res, err1 := T.a.GetS()
	sure.Must(err1)
	return res
}

type Example88Soft struct{ a *Example }

func (a *Example) Soft() *Example88Soft {
	return &Example88Soft{a: a}
}
func (T *Example88Soft) GetN() (res int) {
	res, err1 := T.a.GetN()
	sure.Soft(err1)
	return res
}
func (T *Example88Soft) GetS() (res string) {
	res, err1 := T.a.GetS()
	sure.Soft(err1)
	return res
}
