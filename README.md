# tool

## install 

```bash
go get -u github.com/fuxingZhang/tool
```

## function lists

```go
func Retry(attempts int, sleep time.Duration, fn func() error) error
func Download(url, filepath string) error
func DownloadWithReturnMD5(url, filepath string) (md5Val string, err error) 
func GetFuncName(fn interface{}) string 
func GetFuncNameWithPath(fn interface{}) string 
func PrettyPrint(v interface{}) 
func TrimStruct(src interface{}) error 
func CopyStruct(src, dst interface{}) error
func GetFileMd5(path string) (md5Val string, err error) 
func CheckFileExists(filepath string) bool 
func RandInt64(min, max int64) (num int64, err error)
func RandomStr(n int) string 
```

## usage  

```bash
package main

import (
    "errors"
    "fmt"

    "github.com/fuxingZhang/tool"
)

func main() {
    tool.Retry(3, 1, func() error {
        fmt.Println("----")
        return errors.New("test")
    })

    tool.Download("http://localhost:8080/public/test.sh", "./test.sh")
}
```
