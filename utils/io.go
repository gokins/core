package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5String(data string) string {
	m5 := md5.New()
	m5.Write([]byte(data))
	md5Data := m5.Sum([]byte(nil))
	return hex.EncodeToString(md5Data)
}

func RepSeparators(pth string) string {
	return strings.ReplaceAll(pth, "\\", "/")
}
