package tool

import "testing"

func TestDownload(t *testing.T) {
	err := Download("http://localhost:8080/public/test.sh", "./test.sh")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDownloadAndReturnMD5(t *testing.T) {
	md5, err := DownloadWithReturnMD5("https://alifei04.cfp.cn/creative/vcg/800/version23/VCG41175510742.jpg", "./test.jpg")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(md5)
}
