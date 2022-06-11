package tool

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"net/http"
	"os"
)

// Download download file
func Download(url, filepath string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return
}

// DownloadWithReturnMD5 download file and return md5 string
func DownloadWithReturnMD5(url, filepath string) (md5Val string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		return
	}

	out, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer out.Close()

	h := md5.New()
	writers := io.MultiWriter(out, h)

	_, err = io.Copy(writers, resp.Body)

	md5Val = hex.EncodeToString(h.Sum(nil))
	return
}
