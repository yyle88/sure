package example1

import "github.com/yyle88/mustdone"

type Example88Must struct{ a *Example }

func (a *Example) Must() *Example88Must {
	return &Example88Must{a: a}
}
func (T *Example88Must) GetN() (res int) {
	res, err1 := T.a.GetN()
	mustdone.Must(err1)
	return res
}
func (T *Example88Must) GetS() (res string) {
	res, err1 := T.a.GetS()
	mustdone.Must(err1)
	return res
}

type Example88Soft struct{ a *Example }

func (a *Example) Soft() *Example88Soft {
	return &Example88Soft{a: a}
}
func (T *Example88Soft) GetN() (res int) {
	res, err1 := T.a.GetN()
	mustdone.Soft(err1)
	return res
}
func (T *Example88Soft) GetS() (res string) {
	res, err1 := T.a.GetS()
	mustdone.Soft(err1)
	return res
}
