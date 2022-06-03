# tool

## install 

```bash
go get -u github.com/fuxingZhang/tool
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
}
```