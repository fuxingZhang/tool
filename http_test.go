package tool

import "testing"

func TestDownload(t *testing.T) {
	err := Download("http://localhost:8080/public/test.sh", "./test.sh")
	if err != nil {
		t.Fatal(err)
	}
}
