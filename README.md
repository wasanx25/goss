# goss

goss is text viewer API for Go.

## Use

```go
package main


import (
	"fmt"
	"os"

	"github.com/wasanx25/goss"
)

func main() {
	err := goss.Run("goss is text viewer")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
```
