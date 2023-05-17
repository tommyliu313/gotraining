package main

import (
	"fmt"
	"os"
)

func OpenFile(filename string) (*os.File, error) {
	fmt.Println("Opening", filename)
	return os.Open()
}
