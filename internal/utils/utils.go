package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"unicode"

	"github.com/pkg/errors"
	"github.com/yyle88/done"
	"github.com/yyle88/zaplog"
)

func StringOK(s string) string {
	if s == "" {
		zaplog.ZAPS.P1.LOG.Panic("S IS EMPTY")
	}
	return s
}

func BooleanOK(v bool) bool {
	if !v {
		zaplog.ZAPS.P1.LOG.Panic("B IS FALSE")
	}
	return v
}

func Neat(v interface{}) string {
	data, err := NeatBytes(v)
	if err != nil {
		panic(errors.WithMessage(err, "wrong"))
	}
	return string(data)
}

func NeatBytes(v interface{}) ([]byte, error) {
	data, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, errors.WithMessage(err, "wrong")
	}
	return data, nil
}

func SOrX(s, x string) string {
	if s == "" {
		return x
	}
	return s
}

func In[V comparable](a V, slice []V) bool {
	for _, s := range slice {
		if a == s {
			return true
		}
	}
	return false
}

func RootMustIsExist(root string) bool {
	info, err := os.Stat(root)
	return !os.IsNotExist(err) && info != nil && info.IsDir()
}

func MustWriteToPath(path string, s string) {
	done.Done(os.MkdirAll(filepath.Dir(path), 0755))
	done.Done(os.WriteFile(path, []byte(s), 0644))
}

func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func MustLs(root string) (names []string) {
	infos := done.VAE(os.ReadDir(root)).Done()
	names = make([]string, 0, len(infos))
	for _, info := range infos {
		names = append(names, info.Name())
	}
	return
}

func C0IsUPPER(s string) bool {
	runes := []rune(s)
	if len(runes) > 0 {
		return unicode.IsUpper(runes[0])
	}
	return false
}

func SetDoubleQuotes(s string) string {
	return "\"" + s + "\""
}
