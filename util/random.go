package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

const alphabet = "abcdefghijklmnopqrstuvwz"

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	alphabetLen := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(alphabetLen)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"EURO", "USD", "CAD", "NGN"}
	currencyLen := len(currencies)
	return currencies[rand.Intn(currencyLen)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@mailtest.com", RandomString(6))
}

// RandomOTP generates a random OTP
func RandomOTP() string {
	return fmt.Sprintf("%06d", RandomInt(0, 999999))
}
