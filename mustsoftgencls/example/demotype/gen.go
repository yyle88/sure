package demotype

import (
	"github.com/yyle88/mustdone"
)

type Demo88Must struct{ a *Demo }

func (a *Demo) Must() *Demo88Must {
	return &Demo88Must{a: a}
}
func (T *Demo88Must) GetInt() (res int) {
	res, err1 := T.a.GetInt()
	mustdone.Must(err1)
	return res
}
func (T *Demo88Must) GetFloat64() (res float64) {
	res, err1 := T.a.GetFloat64()
	mustdone.Must(err1)
	return res
}

type Demo88Soft struct{ a *Demo }

func (a *Demo) Soft() *Demo88Soft {
	return &Demo88Soft{a: a}
}
func (T *Demo88Soft) GetInt() (res int) {
	res, err1 := T.a.GetInt()
	mustdone.Soft(err1)
	return res
}
func (T *Demo88Soft) GetFloat64() (res float64) {
	res, err1 := T.a.GetFloat64()
	mustdone.Soft(err1)
	return res
}
