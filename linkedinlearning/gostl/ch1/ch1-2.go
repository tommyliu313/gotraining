package main

import "fmt"

func main() {
	fmt.Print("Welcome to Go!\n")
	fmt.Println("This string ends with a new line")

	const answer = 42
	fmt.Println("The answer to life id", answer)

	const a, b, c = 5, 5, 10
	fmt.Println("Add", a, "and", b, "you get", c)

	items := []int{10, 20, 40, 80}
	length, err := fmt.Println(items)
	fmt.Println(length, err)

}
