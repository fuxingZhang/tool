package tool

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// RandInt Generate random int64 of specified length
func RandInt64(min, max int64) (num int64, err error) {
	n, err := rand.Int(rand.Reader, big.NewInt(max-min))
	if err != nil {
		return
	}
	num = min + n.Int64()
	return
}

// RandomStr Generate random string of specified length
func RandomStr(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}
