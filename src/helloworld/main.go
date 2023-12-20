package main

import (
	"os"

	"github.com/ken1009us/go-practice/cmd/helloworld/hello"
)

func main() {
	// if len(os.Args) > 1 {
	// 	fmt.Println("Hello,", os.Args[1])
	// } else {
	// 	fmt.Println("Hello World!")
	// }

	if len(os.Args) > 1 {
		hello.Say(os.Args[1])
	}
}