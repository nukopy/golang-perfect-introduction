package mapp

import (
	"fmt"
	"strings"
)

func PracticeMap() {
	fmt.Println("===== in chap03/mapp =====")
	// init_map()
	// manipulate_map()
	word_count()
}

func init_map() {
	var mp map[string]int
	fmt.Println(mp)

	// make で初期化
	mp1 := make(map[string]int)
	fmt.Println(mp1)

	// make で初期化。容量を指定できる。
	mp2 := make(map[string]int, 10)
	fmt.Println(mp2)

	// リテラルで初期化
	mp3 := map[string]int{
		"x": 10,
		"y": 20,
	}
	fmt.Println(mp3)

	// リテラルで初期化：空のマップリテラルの場合
	mp4 := map[string]int{}
	fmt.Println(mp4)
}

func manipulate_map() {
	m := map[string]int{"x": 10, "y": 20, "z": 30}

	// 存在するキーにアクセス
	fmt.Println(m["z"])

	// 存在しないキーにアクセス：ゼロ値が返る。ok で存在を確かめられる。
	val, ok := m["xx"]
	fmt.Println(val, ok)
	fmt.Println(m["xx"]) // 存在しないキーにアクセスするとゼロ値が返る。
	val2, ok2 := m["xx"]
	fmt.Println(val2, ok2) // アクセスしたあとも変わらない

	if val, ok := m["xx"]; ok {
		fmt.Println("Key \"xx\" is exists.")
		fmt.Printf("Value: %v\n", val)
	} else {
		fmt.Println("Key \"xx\" is NOT exists.")
		fmt.Printf("Value: %v\n", val)
	}

	// キーを指定して入力
	m["xx"] = 0 // ゼロ値を明示的に代入すると ok が true

	if val, ok := m["xx"]; ok {
		fmt.Println("Key \"xx\" is exists.")
		fmt.Printf("Value: %v\n", val)
	} else {
		fmt.Println("Key \"xx\" is NOT exists.")
		fmt.Printf("Value: %v\n", val)
	}

	// キーを指定して削除
	val3, ok3 := m["xx"]
	fmt.Println("before:", val3, ok3) // 削除前のキーと値
	delete(m, "xx")
	val4, ok4 := m["xx"]
	fmt.Println("after :", val4, ok4) // 削除されていることを確認
}

func word_count() {
	// var counts map[string]int // これだと counts == nil で空のマップですら無く、ランタイムエラーになる
	counts := map[string]int{}
	str := "dog dog dogs cat cat dog cats cats dog dog cat cat"

	for _, s := range strings.Split(str, " ") {
		counts[s]++
	}
	fmt.Println(counts)
}