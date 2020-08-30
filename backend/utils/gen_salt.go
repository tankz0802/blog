package utils

import (
	"math/rand"
	"time"
)

func GenSalt() (salt []byte) {
	rand.Seed(time.Now().Unix())
	salt = make([]byte, 8)
	for i:=0;i<8;i++ {
		salt[i] = byte(rand.Intn(96)) + 32
	}
	return salt
}
