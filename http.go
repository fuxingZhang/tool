package tool

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"time"
)

// Download download file
func Download(url, filepath string) (err error) {
	url, err = EncodeURI(url)
	if err != nil {
		return
	}

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
	Headers      map[string]string
	Timeout      time.Duration
	IsReturnMD5  bool
	IsWithCancel bool
}

// DownloadFileResult dowload file result
type DownloadFileResult struct {
	Err    error
	Md5    string
	Cancel context.CancelFunc
}

// DownloadWithOptions download file with headers and return md5 string
func DownloadWithOptions(url, filepath string, option DownloadFileOption) (result DownloadFileResult) {
	client := &http.Client{
		Timeout: option.Timeout,
	}

	var err error

	defer func() {
		result.Err = err
	}()

	url, err = EncodeURI(url)
	if err != nil {
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	// req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return
	}

	// cancel
	if option.IsWithCancel {
		ctx, cancel := context.WithCancel(context.Background())
		req = req.WithContext(ctx)
		result.Cancel = cancel
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

// EncodeURI Encode URI
func EncodeURI(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return rawURL, err
	}
	u.RawQuery = u.Query().Encode()
	return u.String(), nil
}

// GetFilenameFromHeader get filename from header
func GetFilenameFromHeader(header http.Header) (string, error) {
	contentDisposition := header.Get("Content-Disposition")
	if contentDisposition == "" {
		return "", errors.New("Content-Disposition header not found")
	}
	_, params, err := mime.ParseMediaType("attachment;filename=gateway.doc")
	if err != nil {
		return "", err
	}
	filename, ok := params["filename"]
	if !ok {
		return "", errors.New("filename parameter not found in Content-Disposition header")
	}
	return filename, nil
}
