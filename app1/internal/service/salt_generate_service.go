package service

import (
	"math/rand"
	"time"
)

func GenerateSalt() string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	salt := make([]byte, 12)
	rand.Seed(time.Now().Unix())
	for i := range salt {
		salt[i] = str[rand.Intn(len(str))]
	}
	return string(salt)
}
