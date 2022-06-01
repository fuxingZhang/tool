package tool

import (
	"log"
	"time"
)

// Retry retry function
func Retry(attempts int, sleep time.Duration, fn func() error) (err error) {
	funcName := GetFuncNameWithPath(fn)
	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return
		}
		log.Println(funcName, "run error:", err.Error(), ", retry times:", i)
		time.Sleep(sleep)
	}
	log.Println(funcName, "run failed")
	return
}
