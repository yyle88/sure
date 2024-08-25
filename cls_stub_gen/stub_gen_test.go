package cls_stub_gen

import (
	"testing"

	"github.com/yyle88/runpath"
)

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

func TestGen(t *testing.T) {
	res := GenerateStubFunctions(&Config{
		SrcRoot:       runpath.PARENT.Path(),
		TargetPkgName: "p_p_p",
		ImportOptions: nil,
		TargetSrcPath: "",
		CanCreateFile: false,
	}, &Param{
		object: A{},
		opStub: "a_a_a",
	})
	t.Log(res)
}
