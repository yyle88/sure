package example5sure_custom

import (
	"testing"

	"github.com/pkg/errors"
)

func TestNode_Must(t *testing.T) {
	run := func() (string, error) {
		return "OK", nil
	}
	res, err := run()
	NODE.Must(err)
	t.Log(res)
}

func TestNode_Soft(t *testing.T) {
	run := func() (string, error) {
		return "", errors.New("wrong")
	}
	res, err := run()
	NODE.Soft(err)
	t.Log(res)
}
