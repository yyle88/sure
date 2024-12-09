package cls_stub_gen

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/yyle88/runpath"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

type A struct{}

func (a *A) Get() *A {
	return a
}

func (a *A) Set(string) error {
	return errors.New("not implement")
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

func TestGenerateMethodsStub(t *testing.T) {
	res := GenerateStubMethods(
		&StubGenConfig{
			SourceRootPath:    runpath.PARENT.Path(),
			TargetPackageName: "p_p_p",
			ImportOptions:     syntaxgo_ast.NewPackageImportOptions(),
			OutputPath:        "",
			AllowFileCreation: false,
		},
		NewStubParam(A{}, "singletonInstance"),
	)
	t.Log(res)
}
