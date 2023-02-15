package service

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func GenerateHash(password, salt string) string {
	var str strings.Builder
	str.Grow(len(password) + len(salt))
	str.WriteString(password)
	str.WriteString(salt)
	hash := md5.Sum([]byte(str.String()))
	return hex.EncodeToString(hash[:])
}
