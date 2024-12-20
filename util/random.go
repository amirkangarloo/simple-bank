package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// To generate random owner name
func RandomOwner() string {
	return randomString(6)
}

// To generate random amount money
func RandomMoney() int64 {
	return randomInt(0, 1000)
}

// To generate random currency
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "IRR"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}
