package example4

import "github.com/yyle88/mustdone"

type ExampleMust struct{ a *Example }

func (a *Example) Must() *ExampleMust {
	return &ExampleMust{a: a}
}
func (T *ExampleMust) GetN() (res int) {
	res, err1 := T.a.GetN()
	mustdone.Must(err1)
	return res
}
func (T *ExampleMust) GetS() (res string) {
	res, err1 := T.a.GetS()
	mustdone.Must(err1)
	return res
}

type ExampleSoft struct{ a *Example }

func (a *Example) Soft() *ExampleSoft {
	return &ExampleSoft{a: a}
}
func (T *ExampleSoft) GetN() (res int) {
	res, err1 := T.a.GetN()
	mustdone.Soft(err1)
	return res
}
func (T *ExampleSoft) GetS() (res string) {
	res, err1 := T.a.GetS()
	mustdone.Soft(err1)
	return res
}

type DemoMust struct{ a *Demo }

func (a *Demo) Must() *DemoMust {
	return &DemoMust{a: a}
}
func (T *DemoMust) GetN() (res int) {
	res, err1 := T.a.GetN()
	mustdone.Must(err1)
	return res
}
func (T *DemoMust) GetS() (res string) {
	res, err1 := T.a.GetS()
	mustdone.Must(err1)
	return res
}

type DemoSoft struct{ a *Demo }

func (a *Demo) Soft() *DemoSoft {
	return &DemoSoft{a: a}
}
func (T *DemoSoft) GetN() (res int) {
	res, err1 := T.a.GetN()
	mustdone.Soft(err1)
	return res
}
func (T *DemoSoft) GetS() (res string) {
	res, err1 := T.a.GetS()
	mustdone.Soft(err1)
	return res
}
