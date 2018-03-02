package marketdata

import (
	"math/rand"
	"os"
	"time"
)

// Getwd get working directory
func Getwd() string {
	wd, _ := os.Getwd()
	return wd
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStringRunes generate random string as reqID
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
