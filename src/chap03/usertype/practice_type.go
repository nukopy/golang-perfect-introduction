package usertype

import (
	"fmt"
	"net/http"
	"time"
)

func PracticeType() {
	fmt.Println("===== in chap03/usertype =====")
	// practice_type()
	// user_defined_type1()
	// user_defined_type2()
	TRY_chap03_usertype()
}

func practice_type() {
	// 型として、文法として成り立つかを考える
	var ns []int // ok
	var nss [][]int // ok
	// ...
	var nsss [][][][][][][][][][][][][][][]int

	fmt.Println(ns)
	fmt.Println(ns)
	fmt.Println(nss, len(nss), cap(nss))
	fmt.Println(nsss, len(nsss), cap(nsss))

	// int 型で成り立っていることはコンポジット型で当然できる
}

func user_defined_type1() {
	type MyInt int
	var n int = 100
	m := MyInt(n)
	n = int(m)
	// fmt.Println(n + m) // 型の不一致でコンパイルエラー（別の型として定義される）
	// usertype/practice_type.go:33:16: invalid operation: n + m (mismatched types int and MyInt)

	type MyInt2 = int
	var a int = 100
	var b MyInt2 = 200
	fmt.Println(a + b) // 型エイリアスは同じ型を表すのでコンパイルが通る

	duration := 10 * time.Second
	fmt.Printf("%T: %v\n", duration, duration)
}

func user_defined_type2() {
	type Applicant = http.Client

	fmt.Printf("http.Client{}: %T\nApplicant{}: %T\n", http.Client{}, Applicant{})
}

func TRY_chap03_usertype() {
	type Result struct {
		GameId int
		Name string
		Score int
	}
}