package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
	"unicode"

	"github.com/kr/pretty"
	"github.com/pkg/errors"
	"github.com/yyle88/done"
	"github.com/yyle88/erero"
	"github.com/yyle88/zaplog"
)

func Neat2json(v interface{}) string {
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

func In[V comparable](v V, slice []V) bool {
	return slices.Contains(slice, v)
}

func MustRoot(root string) {
	done.VBE(IsRootExist(root)).TRUE()
}

func MustFile(path string) {
	done.VBE(IsFileExist(path)).TRUE()
}

func IsRootExist(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil // 路径不存在
		}
		return false, erero.Wro(err) // 其他的错误
	}
	return info.IsDir(), nil
}

func IsFileExist(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil // 路径不存在
		}
		return false, erero.Wro(err) // 其他的错误
	}
	return !info.IsDir(), nil
}

func MustWriteIntoPath(path string, s string) {
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

func C0IsUppercase(s string) bool {
	runes := []rune(s)
	if len(runes) > 0 {
		return unicode.IsUpper(runes[0])
	}
	return false
}

func SetDoubleQuotes(s string) string {
	return "\"" + s + "\""
}

func Boolean(v bool) bool {
	return v
}

func PrintObject(a any) {
	zaplog.LOGS.Skip1.Debug("------------")
	zaplog.LOGS.Skip1.Debug("---object---")
	done.VNE(pretty.Println(a)).Fine()
	zaplog.LOGS.Skip1.Debug("------------")
	zaplog.LOGS.Skip1.Debug("------------")
}
