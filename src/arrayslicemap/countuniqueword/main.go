package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "Unique words")

	type kv struct {
		key string
		val int
	}

	var ss []kv

	for k, v := range words {
		ss = append(ss, kv{k ,v})
	}

	// fmt.Println(ss)

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	for _, s := range ss {
		fmt.Println(s.key, "appears", s.val, "times")

	}
}

// We can use scan.Scan() without explicitly setting
// a split function using scan.Split(bufio.ScanWords) first.
// When you create a new bufio.Scanner and don't set
// a specific split function, the scanner uses a default split function.
// In Go's bufio package, the default split function for a scanner
// is bufio.ScanLines.
