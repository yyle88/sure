package example

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/mustdone/mustsoftgencls/example/demotype"
)

func TestExample(t *testing.T) {
	a := demotype.NewDemo(88)
	{
		value, err := a.GetInt()
		require.NoError(t, err)
		t.Log(value)
	}
	{
		value, err := a.GetFloat64()
		require.NoError(t, err)
		t.Log(value)
	}
	t.Log(a.Must().GetInt())
	t.Log(a.Must().GetFloat64())

	t.Log(a.Soft().GetInt())
	t.Log(a.Soft().GetFloat64())
}
