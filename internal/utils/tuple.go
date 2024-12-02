package utils

import "go/ast"

type SourceMethodsTuple struct {
	SourceCode []byte
	MethodList []*ast.FuncDecl
}

type SourceMethodsTuples []*SourceMethodsTuple

func (mts SourceMethodsTuples) GetReceiverVariableName() string {
	for _, codePage := range mts {
		for _, mebFunction := range codePage.MethodList {
			if mebFunction.Recv == nil {
				continue
			}
			if len(mebFunction.Recv.List) == 0 {
				continue
			}
			if len(mebFunction.Recv.List[0].Names) == 0 {
				continue
			}
			name := mebFunction.Recv.List[0].Names[0].Name
			if name != "" {
				return name
			}
		}
	}
	return ""
}
