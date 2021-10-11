package method

import "fmt"

func PracticeMethod() {
	fmt.Println("===== in chap03/method =====")
	// basic()
	// ex1()
	// ex2()
	// TRY1()
	// TRY_me()
	// practice_method_value()
	practice_method_stmt()
}

type Hex int
func (h Hex) String() string {
	return fmt.Sprintf("%d", int(h))
}

func (h *Hex) StringR() string {
	return fmt.Sprintf("%d", int(*h))
}

func basic() {
	// basic
	var hex Hex = 100
	hex_str := hex.String()
	fmt.Printf("%T: %d\n", hex, hex)
	fmt.Printf("%T: %s\n", hex_str, hex_str)
}

func ex1() {
	var hex Hex = 100
	hex_str := hex.String()
	hex_strr := hex.StringR()
	fmt.Printf("%T: %s\n", hex_str, hex_str)
	fmt.Printf("%T: %s\n", hex_strr, hex_strr)
}

type T struct {}
func (t T) f() {
	fmt.Println("f() executed.")
}
func (t *T) g() {
	fmt.Println("g() executed.")
}

// `*T` 型は、`T` のメソッドも自身のメソッドとして扱われる
func ex2() {
	// シンタックスシュガー
	var t T
	(&t).g()
	t.g() // シンタックスシュガー
	// 型 T があたかも g をメソッドとして持っているかのように
	// 扱えるシンタックスシュガー

	// f() はレシーバが T 型
	(T{}).f() // T
	(&T{}).f() // *T
	(*&T{}).f() // T

	// g() はレシーバが *T 型
	// (T{}).g() // T // error: コンパイルエラー
	(&T{}).g() // *T
	(*&T{}).g() // T // 一度ポインタを通すと呼べるようになる
}

type MyInt int
func (n *MyInt) Inc() { *n++ }

func TRY1() {
	var n MyInt
	fmt.Println(n)
	n.Inc()
	fmt.Println(n)
}

// メソッドが呼ばれた回数数えるやつ、ゲッター、セッター
type Counter struct {
	id int
	id_str string
	num_called int
}

func (c *Counter) GetId() int {
	c.num_called++
	return c.id
}
func (c *Counter) GetIdStr() string {
	c.num_called++
	return c.id_str
}
func (c *Counter) SetId(id int) {
	c.num_called++
	c.id = id
}
func (c *Counter) SetIdStr(id_str string) {
	c.num_called++
	c.id_str = id_str
}
func (c *Counter) GetNumCalled() int {
	return c.num_called
}

func TRY_me() {
	c := Counter{id: 1, id_str: "1"} // num_called はゼロ値で 0 が入る
	fmt.Printf("%+v\n", c)
	fmt.Printf("num_called: %d\n", c.GetNumCalled())

	fmt.Printf("id     : %d\n", c.GetId())
	fmt.Printf("id_str : %s\n", c.GetIdStr())
	fmt.Printf("num_called: %d\n", c.GetNumCalled())

	c.SetId(555)
	fmt.Printf("id     : %d\n", c.GetId())
	c.SetIdStr("hogehoge")
	fmt.Printf("id_str : %s\n", c.GetIdStr())
	fmt.Printf("id     : %d\n", c.GetId())
	fmt.Printf("num_called: %d\n", c.GetNumCalled())
}

// メソッド値、メソッド式
type MyHex int
func (h MyHex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func practice_method_value() {
	// メソッド値：メソッド自体を別の変数の値として扱える
	var hex MyHex = 100
	f := hex.String // hex の値は f の宣言時の値 100 に束縛される
	fmt.Println(f())
	hex = 1000
	fmt.Println(f()) // hex の値は束縛されているので、f() の hex の値は変わらず、100 が出力される
}

func practice_method_stmt() {
	// メソッド式：レシーバを第 1 引数とした関数になる
	// 型に紐付いたメソッドを直接呼び出すこともできる
	// また、別の変数の値として渡すことも出来る。その場合、ただの関数のように振る舞う。
	var hex MyHex = 100
	fmt.Println(MyHex.String(hex))

	f := MyHex.String
	// 普通の関数のように振る舞う
	fmt.Println(f(hex))
	fmt.Println(f(1000))
	fmt.Println(f(10000))
}