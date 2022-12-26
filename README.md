# tool

golang common utils functions

## install 

```bash
go get -u github.com/fuxingZhang/tool
```

## doc

[see here](https://pkg.go.dev/github.com/fuxingZhang/tool)

## function lists

```go
func Retry(attempts int, sleep time.Duration, fn func() error) error
func Download(url, filepath string) error
func DownloadWithOptions(url, filepath string, option DownloadFileOption) (result DownloadFileResult)
func GetFuncName(fn interface{}) string 
func GetFuncNameWithPath(fn interface{}) string 
func PrettyPrint(v interface{}) 
func TrimStruct(src interface{}) error 
func CopyStruct(src, dst interface{}) error
func GetFileMd5(path string) (md5Val string, err error) 
func CheckFileExists(filepath string) (exist bool, err error) 
func CheckDirExists(path string) (exist bool, err error) 
func CheckPathExists(path string) (exist bool, err error)
func RandInt64(min, max int64) (num int64, err error)
func RandomStr(n int) string 
func TailFileBySystemCommand(path string, n int) (data []string, err error)
func TailFile(path string, n int) (data []string, err error) 
```

## usage  

```go
package main

import (
    "errors"
    "fmt"

    "github.com/fuxingZhang/tool"
)

func main() {
    err := tool.Retry(3, 1, func() error {
        return errors.New("test")
    })
    fmt.Println(err)

    err = tool.Download("http://localhost:8080/public/test.sh", "./test.sh")
    fmt.Println(err)
}
```

## sub package

### slice

#### function lists

```go
func ContainsInt(ints []int, val int) bool 
func ContainsStr(strs []string, s string) bool 
```

#### usage  

```go
package main

import (
    "errors"
    "fmt"

    "github.com/fuxingZhang/tool/slice"
)

func main() {
    fmt.Println(slice.ContainsInt([]int{1, 2, 3}, 1))
    fmt.Println(slice.ContainsStr([]string{"a", "b"}, "c"))
}
```