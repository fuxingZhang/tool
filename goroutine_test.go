package utils

import (
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	Go(func() {
		panic("test")
	})
	time.Sleep(time.Microsecond * 200)
}
