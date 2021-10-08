package main

import (
	"chap02/omikuji"
	"fmt"
)

func main() {
	// chap02_3()
	chap02_4()
	iota_no_test()
}

func chap02_3() {
	// sum := 0
	var sum int  // sum := 0 と同じ意味を持つ
	fmt.Println(sum)  // sum はゼロ値で初期化される

	for i := 0; i < 10; i++ {
		sum += 1
	}
	// fmt.Printf("sum = %d", sum)
	fmt.Println("ポインタ：", &sum)
	fmt.Println("ポインタが指す値：", *(&sum))
	fmt.Println(sum)
}

func chap02_4() {
	// ------------------------------
	// TRY：奇数と偶数
	// ------------------------------

	// if-else 文を用いる
	// print_odds_or_even_if_stmt()

	// switch 文 version
	// print_odds_or_even_switch_stmt()

	// ------------------------------
	// TRY：おみくじプログラムを作ろう
	// ------------------------------
	fmt.Println(omikuji.Omikuji(omikuji.ThrowDice()))
}

func print_odds_or_even_if_stmt() {
	for i := 1; i <= 100; i++ {
		var odd_or_even string

		if i % 2 != 0 {
			odd_or_even = "奇数"
		} else {
			odd_or_even = "偶数"
		}
		fmt.Printf("%d-%s\n", i, odd_or_even)
	}
}

func print_odds_or_even_switch_stmt() {
	for i := 1; i <= 100; i++ {
		var odd_or_even string

		switch {
		case i % 2 != 0:
			odd_or_even = "奇数"
		default:
			odd_or_even = "偶数"
		}

		fmt.Printf("%d-%s\n", i, odd_or_even)
	}
}
