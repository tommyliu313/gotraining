package main

import (
	"fmt"
	"strings"
)

func main() {
	fname := "filename.txt"
	fname2 := "temp_picfile.jpeg"
	vowels := "aeiouAEIOU"
	s := "The quick brown fox jumps over the lazy dog"

	fmt.Println(strings.Contains(s, "jump"))
	fmt.Println(strings.ContainsAny(s, "abcd"))

	fmt.Println(strings.Index(s, "fox"))
	fmt.Println(strings.Index(s, "cat"))

	fmt.Println(strings.IndexAny("grzbl", vowels))
	fmt.Println(strings.IndexAny("Golang!", vowels))

	fmt.Println(strings.HasSuffix(fname, "txt"))
	fmt.Println(strings.HasPrefix(fname2, "temp"))

	fmt.Println(strings.Count(s, "the"))
	fmt.Println(strings.Count(s, "he"))

}
