package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return rand.Int63n(max-min+1) + min
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyz")

	var sb strings.Builder
	k := len(letter)

	for i := 0; i < n; i++ {
		c := letter[rand.Intn(k)]
		sb.WriteRune(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
