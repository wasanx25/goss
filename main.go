package main

import (
	"fmt"
	"os"

	"github.com/wasanx25/goss/run"
)

func main() {
	data := `
# 見出し1
aa aa aa
## 見出し2
bb	bb	bb
### 見出し3
#### 見出し4


##### 見出し5
`
	err := run.Exec(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
