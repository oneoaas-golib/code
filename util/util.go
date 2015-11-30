package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func CreateDir(dir string) error {
	return os.Mkdir(dir, 0755)
}

func Md5(str string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, str)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsExist(err) {
			return true
		}
	}
	return false
}

func CreateFile(name string) error {
	_, err := os.Create(name)
	return err
}
