# tool

## install 

```bash
go get -u github.com/fuxingZhang/tool
```

## function lists

```go
func Retry(attempts int, sleep time.Duration, fn func() error) error
func Download(url, filepath string) error
func GetFuncName(fn interface{}) string 
func GetFuncNameWithPath(fn interface{}) string 
func PrettyPrint(v interface{}) 
func TrimStruct(src interface{}) error 
func CopyStruct(src, dst interface{}) error
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
