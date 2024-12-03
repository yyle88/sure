package example_cls_stub_gen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExample0x1(t *testing.T) {
	t.Log(Example0x1())
}

func TestExample0x2(t *testing.T) {
	t.Log(Example0x2())
}

func TestExample6x1(t *testing.T) {
	require.Equal(t, "100", Example6x1())
}

func TestExample6x2(t *testing.T) {
	require.Equal(t, "200", Example6x2())
}
