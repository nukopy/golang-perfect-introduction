package main

import (
	"fmt"

	"chap04/greeting"
	"chap04/scope"
)

func main() {
	fmt.Println("===== in chap04 =====")

	// 4.1. パッケージ
	// TRY パッケージを分けてみよう
	greeting.Do()
	// greeting.do() unexported なメソッドなので外部からはアクセスできず、コンパイルエラーになる

	// 4.2. パッケージ変数とスコープ
	scope.CheckUniversalScope()

	// 4.3. パッケージの初期化
	// todo

	// 4.4. ライブラリのバージョン管理
	// todo
}
