package print

import "fmt"

const ARRAY_NUMBER int = 5

func PrintArrayInt(arr [ARRAY_NUMBER]int) {
	fmt.Printf("arr = [\n")
	for i, v := range arr {
		fmt.Printf("\t%d: %d,\n", i, v)
	}
	fmt.Printf("]\n")
}

func PrintArrayString(arr [ARRAY_NUMBER]string) {
	fmt.Printf("arr = [\n")
	for i, v := range arr {
		fmt.Printf("\t%d: %s,\n", i, v)
	}
	fmt.Printf("]\n")
}

func PrintArrayBool(arr [ARRAY_NUMBER]bool) {
	fmt.Printf("arr = [\n")
	for i, v := range arr {
		fmt.Printf("\t%d: %t,\n", i, v)
	}
	fmt.Printf("]\n")
}