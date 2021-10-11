package main

import "fmt"

func main() {
	fmt.Println("===== in chap04 =====")
	// tmp()
}

// ------------------------------
// tmp: for ループの挙動を調べた
// ------------------------------

type Person struct {
	Name string
	BirthYear int
}

func (p *Person) Greet() {
	fmt.Printf("Hello, my name is %s!\n", p.Name)
	// fmt.Printf("Hello, my name is %s!\n", (*p).Name)
}

func tmp() {
	ps := []Person{
		{"Bob", 1990},
		{"Alice", 1985},
		{"Carol", 1995},
	}
	fs := make([]func(), 0, len(ps))

	for _, p := range ps {
		fmt.Printf("%p, %p\n", &p, (&p).Greet)
		p.Greet()

		fs = append(fs, p.Greet)
	}

	fmt.Println(fs[0])
	fmt.Println(fs[1])
	fmt.Println(fs[2])

	fmt.Printf("----------\n")

	for _, f := range fs {
		f()
	}
}