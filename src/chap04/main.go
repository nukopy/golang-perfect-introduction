package main

import (
	"fmt"
	"time"

	"chap04/greeting"

	tgreeting "github.com/tenntenn/greeting"
	tgreetingv2 "github.com/tenntenn/greeting/v2"
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
	greeting.Do()
	// greeting.do() unexported なメソッドなので外部からはアクセスできず、コンパイルエラーになる

	// 4.2. パッケージ変数とスコープ
	// scope.CheckUniversalScope()

	// 4.3. パッケージの初期化
	// init 関数の実行

	// 4.4. ライブラリのバージョン管理
	// v1
	fmt.Println("===== test github.com/tenntenn/greeting =====")
	tgreeting.Do()
	// v2
	// JST := time.FixedZone("Asia/Tokyo", 9*60*60)
	JST := time.FixedZone("Asia/Tokyo", 9*60*60)
	date1 := time.Date(2020, 10, 31, 1, 0, 0, 0, time.UTC).In(JST)
	date2 := time.Date(2020, 10, 31, 14, 0, 0, 0, time.UTC).In(JST)

	fmt.Printf("date1: %v\n", date1)
	fmt.Println(tgreetingv2.Do(date1))

	fmt.Printf("date2: %v\n", date2)
	fmt.Println(tgreetingv2.Do(date2))
}
