package main

import (
	"chap03/print"
	"chap03/slice"
	"fmt"
)

func main() {
	// type
	// chap03_1_cast()
	// chap03_1_cast_exercise_error()
	// chap03_1_cast_exercise()
	// chap03_1_type_bool_exercise()

	// struct
	// chap03_1_type_literal()
	// chap03_1_struct()

	// array
	// chap03_1_array_init()
	// chap03_1_array_manipulate()

	// composite literals
	// chap03_1_composite_literals()

	// slice
	chap03_1_slice()
}

// ------------------------------
// スライス
// ------------------------------

func chap03_1_slice() {
	slice.PracticeSlice()
}

// ------------------------------
// Composite literals
// ------------------------------

func chap03_1_composite_literals() {
	// 構造体リテラルで初期化
	myStruct := struct {
		name string
		age int
	}{
		name: "Gopher",
		age: 10,
	}
	fmt.Println(myStruct)

	// 構造体リテラルでフィールド名を省略
	myStruct2 := struct {
		name string
		age int
	}{
		"Gopher",
		10,
	}
	fmt.Println(myStruct2)

	// （コンパイルエラー）構造体リテラルでインデックスを使う
	/*
	myStruct3 := struct {
		name string
		age int
	}{
		0: "Gopher",
		1: 10,
	}
	fmt.Println(myStruct3)
	*/
}

// ------------------------------
// 配列
// ------------------------------

func chap03_1_array_manipulate() {
	ns := [...]int{10, 20, 30, 40, 50}
	// 添字は 0 から始まる

	// 要素にアクセス
	println(ns[3]) // 添字は式（評価されて値が返る）で記述する。変数でも、1 + 1 みたいな式でも良い。

	// 配列の長さ
	println(len(ns))

	// for 文で要素を取得
	for i, v := range ns {
		fmt.Printf("%d: %d\n", i, v)
	}

	// スライス演算
	fmt.Println(ns[1:3]) // インデックス「1」から「3 の 1 つ手前」までの要素を取得)
	// [20, 30]
}

func chap03_1_array_init() {
	// ゼロ値で初期化
	// ゼロ値 - int 型: 0、string 型: ""、bool 型: false）
	var (
		nsInt [5]int
		nsString [5]string
		nsBool [5]bool
	)
	print.PrintArrayInt(nsInt)
	print.PrintArrayString(nsString)
	print.PrintArrayBool(nsBool)

	// 配列リテラルで初期化
	ns1 := [5]int{10, 20, 30, 40, 50}
	print.PrintArrayInt(ns1)

	// （コンパイラが）要素数を値から推論
	ns2 := [...]int{10, 20, 30, 40, 50}
	print.PrintArrayInt(ns2)

	// 特定のインデックスの要素だけ指定して初期化
	// 5 番目が 50、10 番目が 100 で他が 0（ゼロ値）の要素数 11 の配列
	ns4 := [...]int{5: 50, 10: 100}
	for i, v := range ns4 {
		fmt.Printf("%d: %d\n", i, v)
	}
}

// ------------------------------
// 構造体
// ------------------------------

func chap03_1_struct() {
	p := struct {
		name string
		age int
	}{
		name: "Gopher",
		age: 10,
	}

	fmt.Println(p.name, p.age)
	p.age++
	fmt.Println(p.name, p.age)
}

// ------------------------------
// 型
// ------------------------------

func chap03_1_cast() {
	var f float64 = 10
	var n int = int(f)
	fmt.Println(n)
}

/*
func chap03_1_cast_exercise_error() {
	var sum int
	sum = 5 + 6 + 3
	avg := sum / 3
	// avg := float32(sum) / 3
	if avg > 4.5 { // compile error
		fmt.Println("good")
	}
}
*/

func chap03_1_cast_exercise() {
	var sum int
	sum = 5 + 6 + 3
	avg := float32(sum) / 3
	if avg > 4.5 {
		fmt.Println(sum)
		fmt.Println("good")
	}
}

func chap03_1_type_bool_exercise()  {
	var a, b, c bool // ゼロ値: false

	if a && b || !c {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func chap03_1_type_literal() {
	// int 型のスライスの形リテラルを使った変数定義
	var ns []int
	var m map[string]int

	fmt.Println(ns)
	fmt.Println(m)
}