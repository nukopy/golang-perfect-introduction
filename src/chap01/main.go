/*
cd src/chap01

# go build（ビルド）
go build main.go
./main

# go run（ビルド & 実行）
go run main.go
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	mymod()
}
