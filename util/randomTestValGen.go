package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)

}
func randomString(n int) string {
	k := len(alphabet)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()

}

func RandomMoney() int64 {
	return randomInt(0, 10000)

}
func RandomOwner() string {
	return randomString(7)
}

func RandomCurrency() string {
	curr := []string{"USD", "NPR", "GBP"}
	n := len(curr)

	return curr[rand.Intn(n)]

}
