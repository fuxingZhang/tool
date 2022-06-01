package utils

import (
	"runtime"

	"github.com/rs/zerolog/log"
)

// Go Go
func Go(cb func()) {
	go func() {
		defer handlePanic()
		cb()
	}()
}

func handlePanic() {
	if r := recover(); r != nil {
		log.Error().Msgf("Recovered in HandlePanic: %v", r)

		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		log.Error().Str("stack", string(buf[:n])).Msg("goroutine panic")
	}
}
