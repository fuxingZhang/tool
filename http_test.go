package tool

import (
	"os"
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
