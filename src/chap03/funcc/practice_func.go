package funcc

import "fmt"

func PracticeFunc() {
	fmt.Println("===== in chap03/funcc =====")
	// practice_closure()
	// practice_closure2()
	// practice_copy()
	// practice_pointer()
	// odd_or_even()
	// practice_swap()
}

func practice_closure() {
	fs := make([]func() int, 3)

	for i := range fs {
		fs[i] = func() int {
			fmt.Println("i:", i)
			return i
		}

		fs[i]()
	}

	fmt.Println("-----")

	for _, f := range fs { f() }
}

func practice_closure2() {
	// 無名関数（クロージャ）
	msg := "Hello, 世界!"
	nonamef := func() {
		fmt.Printf("This is a closure.\n")
		fmt.Println(msg)
	}
	nonamef()
	msg = "Hello, world!"
	nonamef()
}

func practice_copy() {
	p := struct {
		age int
		name string
	}{
		age: 10,
		name: "Bob",
	}

	p2 := p
	p2.age = 35

	fmt.Println(p)
	fmt.Println(p2)
}


func multiply_100(x_ptr *int) {
	// 元の値を 100 倍する
	*x_ptr *= 100
}

func practice_pointer() {
	ns := []int{10, 20, 30}
	ns1 := ns // 同じものを参照している
	ns[1] = 200
	fmt.Println(ns)
	fmt.Println(ns1)

	multiply_100(&ns[2])
	fmt.Println(ns)
	fmt.Println(ns1)
}

func isOdd(n int) bool {
	return n % 2 != 0
}

func odd_or_even() {
	for i := 1; i <= 100; i++ {
		if isOdd(i) {
			fmt.Println(i, "is odd")
		} else {
			fmt.Println(i, "is even")
		}
	}
}

func _swap(n int, m int) (int, int) {
	return m, n
}

func swap(n *int, m *int) {
	*n, *m = *m, *n
}

func practice_swap() {
	n, m := 10, 20
	fmt.Printf("n: %v, m: %v\n", n, m)
	n, m = _swap(n, m) // n, m = m, n で swap できる
	fmt.Printf("swapped\n")
	fmt.Printf("n: %v, m: %v\n", n, m)
	swap(&n, &m)
	fmt.Printf("swapped\n")
	fmt.Printf("n: %v, m: %v\n", n, m)
}
