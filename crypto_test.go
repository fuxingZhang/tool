package tool

import (
	"log"
	"testing"
)

func TestGetFileMd5(t *testing.T) {
	md5, err := GetFileMd5("./crypto.go")
	if err != nil {
		log.Fatal(err)
	}
	t.Log(md5)
}
