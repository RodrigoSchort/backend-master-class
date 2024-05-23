package utils

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	if min > max {
		panic("min cannot be greater than max")
	}

	diff := big.NewInt(max - min + 1)
	r, _ := rand.Int(rand.Reader, diff)

	return min + r.Int64()
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		v, _ := rand.Int(rand.Reader, big.NewInt(int64(k)))

		c := alphabet[v.Int64()]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	owner := RandomString(6)

	return owner
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"BRL", "USD", "EUR"}
	n := len(currencies)

	randIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(n)))

	return currencies[randIndex.Int64()]
}
