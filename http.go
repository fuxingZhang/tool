package tool

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"net/http"
	"os"
	"time"
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

// DownloadFileOption download file option
type DownloadFileOption struct {
	Headers     map[string]string
	Timeout     time.Duration
	IsReturnMD5 bool
}

// DownloadFileResult dowload file result
type DownloadFileResult struct {
	Err error
	Md5 string
}

// DownloadWithOptions download file with headers and return md5 string
func DownloadWithOptions(url, filepath string, option DownloadFileOption) (result DownloadFileResult) {
	client := &http.Client{
		Timeout: option.Timeout,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	// headers
	for k, v := range option.Headers {
		req.Header.Add(k, v)
	}

	// requset
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		return
	}

	// create
	out, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer out.Close()

	if !option.IsReturnMD5 {
		_, err = io.Copy(out, resp.Body)
		return
	}

	// md5
	h := md5.New()
	writers := io.MultiWriter(out, h)

	_, err = io.Copy(writers, resp.Body)

	result.Md5 = hex.EncodeToString(h.Sum(nil))
	return
}
