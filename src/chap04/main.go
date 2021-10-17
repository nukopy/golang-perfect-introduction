package main

import (
	"fmt"
)

func init() {
	// main の前に実行される
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
	main()
	// main 関数は明示的に呼び出すことが出来る
	// ここで呼び出した main とは別にまた main が呼ばれる。
	// init はあくまでパッケージの初期化処理なので、
	// ここで呼び出した main は初期化処理 1 つとして扱われる。
}

func main() {
	fmt.Println("===== in chap04 =====")
	fmt.Println("main")

	// 4.1. パッケージ
	// TRY パッケージを分けてみよう
	// greeting.Do()
	// greeting.do() unexported なメソッドなので外部からはアクセスできず、コンパイルエラーになる

	// 4.2. パッケージ変数とスコープ
	// scope.CheckUniversalScope()

	// 4.3. パッケージの初期化
	// init 関数の実行

	// 4.4. ライブラリのバージョン管理
	// todo
}
