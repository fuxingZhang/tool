package tool

import (
	"io"
	"net/http"
	"os"
)

func Download(url, filepath string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return
}
