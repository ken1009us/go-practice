package main

import (
	"fmt"
)

func main() {
	s := "élite"
	// s := "我是"

	// The length of the string is the number of bytes required
	// to represent the unicode characters,
	// not the number of the unicode characters
	fmt.Printf("%8T %[1]v %d\n", s, len(s))
	fmt.Printf("%8T %[1]v\n", []rune(s))

	b := []byte(s)
	fmt.Printf("%8T %[1]v %d\n", b, len(b))
}