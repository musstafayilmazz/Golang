package util

import (
	"github.com/bxcodec/faker/v4"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

}

// RandomInt generates random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of given length
func RandomString(length int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < length; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return faker.FirstName() + " " + faker.LastName()
}

// RandomMoney generates random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 9999)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "TRY"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
