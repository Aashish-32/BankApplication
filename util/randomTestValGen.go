package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)

}
func RandomString(n int) string {
	k := len(alphabet)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()

}

func RandomMoney() int64 {
	return RandomInt(0, 10000)

}
func RandomOwner() string {
	return RandomString(7)
}

func RandomCurrency() string {
	curr := []string{"USD", "EUR", "GBP"}
	n := len(curr)

	return curr[rand.Intn(n)]

}
func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}
