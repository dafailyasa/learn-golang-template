package tools

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwqyz"

func RandomString(n int) string {
	var sb strings.Builder
	k := big.NewInt(int64(len(alphabet)))

	for i := 0; i < n; i++ {
		randomIndex, err := rand.Int(rand.Reader, k)
		if err != nil {
			// Handle error
			return ""
		}
		c := alphabet[randomIndex.Int64()]
		sb.WriteByte(byte(c))
	}

	return sb.String()
}
