package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Not enough args.")
		os.Exit(-1)
	}

	// Initialize two vars at the same time
	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		// fmt.Println(s)
		t := strings.Join(s, new)
		fmt.Println(t)
	}

}