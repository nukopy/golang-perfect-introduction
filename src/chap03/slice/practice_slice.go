package slice

import "fmt"

func PracticeSlice() {
	fmt.Println("=== in chap03/slice.PracticeSlice() ===")
	// basic()
	// init_slice()
	// comparison_slice_and_array()
	// manipulate_slice()
	// check_address_array_and_slice()
	// practice_append()
	// slice_of_slice()
	TRY_use_of_slice()
}

func PrintSliceInfo(name string, ns []int) {
	fmt.Printf("NAME: %s\n%v\nlen: %d\ncap: %d\n-----\n", name, ns, len(ns), cap(ns))
}

// ------------------------------
// practices
// ------------------------------

func basic() {
	ns := []int{10, 20, 30, 40, 50}
	slice := ns[1:3]

	fmt.Println(ns)
	fmt.Println(slice)
}

func init_slice() {
	// ゼロ値は nil
	var ns1 []int
	PrintSliceInfo("ns1", ns1)

	// 長さと容量を指定して初期化
	// 各要素はゼロ値で初期化されている
	ns1 = make([]int, 3, 10)
	PrintSliceInfo("ns1", ns1)

	// スライスリテラルで初期化
	var ns2 = []int{10, 20, 30, 40, 50}
	PrintSliceInfo("ns2", ns2)

	// 5 番目が 10、10 番目が 100 で他が 0 の要素数 11 のスライス
	ns3 := []int{5: 50, 10: 100}
	PrintSliceInfo("ns3", ns3)
}

func comparison_slice_and_array() {
	// ns と ns_ はだいたい同じ処理
	ns := make([]int, 3, 10)
	PrintSliceInfo("ns", ns)

	var arr [10]int
	ns_ := arr[0:3]
	PrintSliceInfo("ns_", ns_)

	// ms と ms_ はだいたい同じ処理
	ms := []int{10, 20, 30, 40, 50}
	PrintSliceInfo("ms", ms)

	var arr2 = [...]int{10, 20, 30, 40, 50}
	ms_ := arr2[0:5]
	PrintSliceInfo("ms_", ms_)
	// ms_ := arr2[:5]
	// ms_ := arr2[:]
}

func manipulate_slice() {
	ns := []int{10, 20, 30, 40, 50}

	// access to element
	fmt.Println(ns[3])

	// length
	fmt.Println(len(ns))

	// cap
	fmt.Println(cap(ns))
	fmt.Println("----------")

	// append element
	// append 関数は新しいスライスを返す
	// 容量（cap）が足りない場合は背後の配列が再確保される。
	// 前後の len と cap の違いに着目
	PrintSliceInfo("ns", ns)
	ns = append(ns, 60, 70)
	PrintSliceInfo("ns", ns)
}

func check_address_array_and_slice() {
	arr := [...]int{10, 20, 30, 40, 50}
	arr_ptr := &arr[1]
	fmt.Println(arr_ptr)

	ns := arr[1:3]
	ns_ptr := &ns[0]
	fmt.Println(ns_ptr) // arr_ptr と一致する（はず）
	PrintSliceInfo("ns", ns)

	// append して arr に再代入するとメモリの別の場所に確保されるためアドレスは変わる？
	ns = append(ns, 60, 70)
	ns_ptr = &ns[0] // アドレスが変わる？
	fmt.Println(ns_ptr)
	PrintSliceInfo("ns", ns)
}

func practice_append() {
	a := []int{10, 20, 30} // len: 3, cap: 3
	// a: [10 20 30]
	PrintSliceInfo("a", a)

	b := append(a, 40) // len: 4, cap: 6(?)
	// b: [10 20 30 40]
	a[0] = 100
	// a: [100 20 30 40]
	PrintSliceInfo("a", a)
	PrintSliceInfo("b", b)

	c := append(b, 50) // len: 5, cap: 6(?)
	b[1] = 200
	// c: [100 20 30 40 50]
	PrintSliceInfo("b", b)
	PrintSliceInfo("c", c)

	// print address
	fmt.Printf("address of a[0]: %p\n", &a[0])
	fmt.Printf("address of b[0]: %p\n", &b[0])
	fmt.Printf("address of c[0]: %p\n", &c[0])

	// 考え方
	/*
	a に append すると、cap が足りず
	メモリの再確保（値のコピー）が起きるため、
	a[0] と b[0] のアドレスは変わる。

	b に append しても cap は足りているので
	メモリの再確保は起きず、b はそのまま使われる。
	そのため、b[0] と c[0] のアドレスは変わらない。
	そして b も c も同じ背後の配列を参照しているため、
	b[1] = 200 の変更が b、c 両方に反映される。
	*/
}

func slice_of_slice() {
	ns := []int{10, 20, 30, 40, 50}
	n, m := 2, 4

	// n 番目以降のスライスを取得する
	fmt.Println(ns[n:])

	// 先頭から m-1 番目までのスライスを取得する
	fmt.Println(ns[:m])

	PrintSliceInfo("ns", ns)

	// slice の cap を指定する
	// cap 未指定
	ns1 := ns[:2] // len: 2, cap: 5
	PrintSliceInfo("ns1", ns1)

	// cap 指定
	ns2 := ns[:2:3] // len: 2, cap: 3
	PrintSliceInfo("ns2", ns2)

	// もともとのスライスの cap より大きい、また未満を指定してみる
	// cap より大きい
	ns3 := ns[:2:1000] // len: 2, cap: 5
	// ns3 := ns[:2:10] // len: 2, cap: 5
	PrintSliceInfo("ns3", ns3)
	// ランタイムエラーになる
	// 10 でも 1000 でも関係ない。
	// 現在の cap を超えたものをスライスに指定するとランタイムエラーになるらしい。
	/*
	panic: runtime error: slice bounds out of range [::1000] with capacity 5

	goroutine 1 [running]:
	chap03/slice.slice_of_slice()
			/Users/pyteyon/Projects/Languages/Go/golang-perfect-introduction/src/chap03/slice/practice_slice.go:170 +0x154
	chap03/slice.PracticeSlice()
			/Users/pyteyon/Projects/Languages/Go/golang-perfect-introduction/src/chap03/slice/practice_slice.go:13 +0x57
	main.chap03_1_slice(...)
			/Users/pyteyon/Projects/Languages/Go/golang-perfect-introduction/src/chap03/main.go:36
	main.main()
			/Users/pyteyon/Projects/Languages/Go/golang-perfect-introduction/src/chap03/main.go:28 +0x18
	*/

	// 未満（ns の cap は 5 なのでそれより小さい数を指定）
	/*
	ns4 := ns[:2:1] // len: 2, cap: 1
	PrintSliceInfo("ns4", ns4)
	*/

	// コンパイルエラーになる（インデックスが invalid）
	/*
	slice/practice_slice.go:188:11: invalid slice index: 2 > 1
	*/
}

func TRY_use_of_slice_before() {
	// プログラムの動作をそのままに、
	// 3 つの変数しか使わないように修正しろ
	n1 := 19
	n2 := 86
	n3 := 1
	n4 := 12

	sum := n1 + n2 + n3 + n4
	fmt.Println(sum)
}

func TRY_use_of_slice() {
	ns := []int{19, 86, 1, 12}
	var sum int

	for i := range ns {
		sum += ns[i]
	}
	fmt.Println(sum)
}