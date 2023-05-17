package main

import (
	"fmt"
)

func main() {
	s := "The quick brown fox jumps over the lazy dog"

	fmt.Println(len(s))

	for _, ch := range s {
		fmt.Print(string(ch), ",")
	}
}
