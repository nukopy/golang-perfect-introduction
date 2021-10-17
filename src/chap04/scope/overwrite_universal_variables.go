package scope

import (
	"fmt"
)

func CheckUniversalScope() {
    // ユニバーススコープに属する false の書き換え
    fmt.Println("----- Check overwrite \"false\" -----")
    defaultFalse()
    overwriteFalse()

    // ユニバーススコープに属する iota の書き換え
    fmt.Println("----- Check overwrite \"iota\" -----")
    defaultIota()
    overwriteIota()
}

func defaultFalse() {
    if false {
        fmt.Println("false is truthy.")
        fmt.Printf("value: %v\n", false)
    } else {
        fmt.Println("false is falsy.")
        fmt.Printf("value: %v\n", false)
    }
    /* 補足
    関数 overwriteFalse で書き換えた false はブロックスコープ内でのみ有効なので、
    関数 defaultFalse の false はユニバーススコープの false を参照している（いつも使っている false と同じ false）。
    */
}

func overwriteFalse() {
    false := true
    if false {
        fmt.Println("false is truthy.")
        fmt.Printf("value: %v\n", false)
    } else {
        fmt.Println("false is falsy.")
        fmt.Printf("value: %v\n", false)
    }
}

func defaultIota() {
    const (
        a = iota
        b
        c
    )
    fmt.Println(a, b, c)
    /* 補足
    iota はブロックスコープ内でのみ有効なので、
    関数 defaultIota の iota はユニバーススコープの iota を参照している（いつも使っている iota と同じ iota）。
    */
}

func overwriteIota() {
    const iota = 11
    const (
        a = iota
        b
        c
    )

    fmt.Println(a, b, c)
}