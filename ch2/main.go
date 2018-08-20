package main

import "fmt"

func main() {
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello, " + "World")
	fmt.Println("Enter a number: ")
	var input float64

	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
