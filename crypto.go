package tool

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// GetFileMd5 get file md5 string
func GetFileMd5(path string) (md5Val string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}

	defer f.Close()

	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return
	}

	md5Val = hex.EncodeToString(h.Sum(nil))
	return
}
