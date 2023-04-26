package tool

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestDownload(t *testing.T) {
	err := Download("http://localhost:8080/public/test.sh", "./test.sh")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDownloadAndReturnMD5(t *testing.T) {
	path := "./test.jpg"
	res := DownloadWithOptions("https://alifei04.cfp.cn/creative/vcg/800/version23/VCG41175510742.jpg",
		path,
		DownloadFileOption{})
	if res.Err != nil {
		t.Fatal(res.Err)
	}
	t.Log(res.Md5)

	res = DownloadWithOptions("https://alifei04.cfp.cn/creative/vcg/800/version23/VCG41175510742.jpg",
		path,
		DownloadFileOption{IsReturnMD5: true})
	if res.Err != nil {
		t.Fatal(res.Err)
	}
	t.Log(res.Md5)

	os.Remove(path)
}

func TestEncodeURI(t *testing.T) {
	u := "http://192.168.108.24:18080/onecity/bc/boss/minio/download/gateway合并.doc?appCode=oneops-auto-maintainance&objectName=2023/04/24/164655524/gateway合并.doc&type=1"

	{
		url, err := EncodeURI(u)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(url)
	}

	{
		var EncodeURI = func(rawURL string) (string, error) {
			u, err := url.Parse(rawURL)
			if err != nil {
				return rawURL, err
			}
			return fmt.Sprintf("%s://%s%s?%s", u.Scheme, u.Host, u.EscapedPath(), u.Query().Encode()), nil
		}
		v, err := EncodeURI(u)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(v)
	}

	// JS EncodeURI
	// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/encodeURI
	{
		var EncodeURI = func(u string) string {
			u = url.QueryEscape(u)
			u = strings.ReplaceAll(u, "%2A", "*")
			u = strings.ReplaceAll(u, "%2B", "+")
			u = strings.ReplaceAll(u, "%2C", ",")
			u = strings.ReplaceAll(u, "%2F", "/")
			u = strings.ReplaceAll(u, "%3A", ":")
			u = strings.ReplaceAll(u, "%3B", ";")
			u = strings.ReplaceAll(u, "%3D", "=")
			u = strings.ReplaceAll(u, "%3F", "?")
			u = strings.ReplaceAll(u, "%21", "!")
			u = strings.ReplaceAll(u, "%23", "#")
			u = strings.ReplaceAll(u, "%24", "$")
			u = strings.ReplaceAll(u, "%26", "&")
			u = strings.ReplaceAll(u, "%27", "'")
			u = strings.ReplaceAll(u, "%28", "(")
			u = strings.ReplaceAll(u, "%29", ")")
			u = strings.ReplaceAll(u, "%40", "@")
			return u
		}
		v := EncodeURI(u)
		t.Log(v)
	}

	// JS EncodeURI
	// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/encodeURI
	{
		var EncodeURI = func(u string) string {
			u = url.PathEscape(u)
			u = strings.ReplaceAll(u, "%2A", "*")
			// u = strings.ReplaceAll(u, "%2B", "+")
			u = strings.ReplaceAll(u, "%2C", ",")
			u = strings.ReplaceAll(u, "%2F", "/")
			// u = strings.ReplaceAll(u, "%3A", ":")
			u = strings.ReplaceAll(u, "%3B", ";")
			// u = strings.ReplaceAll(u, "%3D", "=")
			u = strings.ReplaceAll(u, "%3F", "?")
			u = strings.ReplaceAll(u, "%21", "!")
			u = strings.ReplaceAll(u, "%23", "#")
			// u = strings.ReplaceAll(u, "%24", "$")
			// u = strings.ReplaceAll(u, "%26", "&")
			u = strings.ReplaceAll(u, "%27", "'")
			u = strings.ReplaceAll(u, "%28", "(")
			u = strings.ReplaceAll(u, "%29", ")")
			// u = strings.ReplaceAll(u, "%40", "@")
			return u
		}
		v := EncodeURI(u)
		t.Log(v)
	}
}

func TestGetFilenameFromHeader(t *testing.T) {
	{
		var header = http.Header{}
		header.Add("Content-Disposition", "attachment;filename=gateway.doc")
		filename, err := GetFilenameFromHeader(header)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(filename)
	}

	{
		resp, err := http.Get("http://cachefly.cachefly.net/100mb.test")
		if err != nil {
			return
		}
		defer resp.Body.Close()
		filename, err := GetFilenameFromHeader(resp.Header)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(filename)
	}
}
