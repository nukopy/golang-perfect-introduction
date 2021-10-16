package greeting

import (
	"fmt"
)

// exported
func Do() {
	fmt.Println("Do: こんにちは")
	do() // unexported だけど内部ではもちろん参照可能
}

// unexported
func do() {
	fmt.Println("do: こんにちは")
}