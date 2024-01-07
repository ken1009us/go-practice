package main

import "fmt"



func main() {
	t := []byte("string")

	fmt.Println(len(t), t) // 6 [115 116 114 105 110 103]
	fmt.Println(t[2]) // 104
	fmt.Println(t[:2]) // [115 116]
	fmt.Println(t[2:]) // [114 105 110 103]
	fmt.Println(t[3:5]) // [105 110]


	// var b [3]int
	// b[0] = 1
	// c := [3]int{0, 0, 0}
	// fmt.Println(b, c)


	// Example
	var w = [...]int{1, 2 ,3}
	var x = []int{0, 0, 0}

	y := do(w, x)
	fmt.Println(w, x, y) // [1 2 3] [3 0 0] [3 0 0 0 42]
	// w the values were copied into a new place - the variable a,
	// and now I'm changing a but not changing w

}

func do(a [3]int, b []int) []int {
	// a = b // SYNTAX ERROR since they are different types
	a[0] = 4 // w unchanged
	b[0] = 3 // x changed

	c := make([]int, 5) // []int{0, 0, 0, 0, 0}
	c[4] = 42
	copy(c, b) // copies only 3 elts

	return c
}

