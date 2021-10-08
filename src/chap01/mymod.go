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

func mymod() {
	var1 :=	"string"
	var2 :=	1

	fmt.Println("Hello, world!")
	fmt.Println("hoge")
	fmt.Println(var1, var2)
}
