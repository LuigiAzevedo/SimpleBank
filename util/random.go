package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // 0 -> max-min
}

// RandomFloat generates a random float64 between min and max
func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min) // 0 -> max-min
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a number string with 2 decimals between min and max
func RandomMoney(min, max float64) string {
	return fmt.Sprintf("%.2f", RandomFloat(min, max))
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"EUR", "USD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
