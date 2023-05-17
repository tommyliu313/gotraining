package main

import "fmt"

type circle struct {
	radius int
	border int
}

func main() {

	x := 20
	f := 123.45

	fmt.Printf("%d\n", x)
	fmt.Printf("%x\n", x)

	fmt.Printf("%t\n", x > 10)

	fmt.Printf("%f\n", f)
	fmt.Printf("%e\n", f)

}
