package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}


//RandomInt generated btw min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

//RandomString generated
func RandomString (n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// generate random owner name
func RandomOwner() string{
	return RandomString(6)
}

// generate random balance
func RandomBalance() int64 {
	return RandomInt(0,1000)
}

//generate random currency
func RandomCurrency() string {
	currencies := []string{"USD", "EURO", "POUND"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}