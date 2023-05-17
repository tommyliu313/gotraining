package main

import "fmt"

func main() {
	f := 123.4567

	fmt.Printf("%.2f\n", f)
	fmt.Printf("%10f\n", f)

	fmt.Printf("%10.2f\n", f)
	fmt.Printf("%+10.2f\n", f)
}
