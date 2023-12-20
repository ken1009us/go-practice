package main

import "strings"

// func SayHello(name string) string {
// 	return fmt.Sprintf("Hello, %s!", name)
// }


func SayHello(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}

	return "Hello, " + strings.Join(names, ", ") + "!"
}