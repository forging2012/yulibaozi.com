package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"golang.org/x/crypto/scrypt"
)

// Encryption 对指定字符串加密（专家级加密方式）
func Encryption(key, salt string) (mdwpd string, err error) {
	var (
		dk []byte
	)
	dk, err = scrypt.Key([]byte(key), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		return
	}
	// []byte-->string
	mdwpd = hex.EncodeToString(dk)
	return
}

// Md5 接口效验
func Md5(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	return strings.ToLower(fmt.Sprintf("%x", m.Sum(nil)))
}
