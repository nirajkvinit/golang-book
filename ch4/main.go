package main

import "fmt"

func main() {
	/*i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}*/
	// for i := 1; i <= 10; i++ {
	// 	/*if i % 2 == 0 {
	// 		fmt.Println(i, "even")
	// 	} else {
	// 		fmt.Println(i, "odd")
	// 	}*/

	// 	switch i {
	// 	case 0:
	// 		fmt.Println("Zero")
	// 	case 1:
	// 		fmt.Println("One")
	// 	case 2:
	// 		fmt.Println("Two")
	// 	case 3:
	// 		fmt.Println("Three")
	// 	case 4:
	// 		fmt.Println("Four")
	// 	case 5:
	// 		fmt.Println("Five")
	// 	case 6:
	// 		fmt.Println("Six")
	// 	default:
	// 		fmt.Println("Unknown Number")
	// 	}
	// }

	/* Follwing doesn't work
	for (i := 1; i <= 10; i++) {
		fmt.Println(i)
	}*/

	/*Write a program that prints out all the numbers between 1 and 100 that are
	evenly divisible by 3 (i.e., 3, 6, 9, etc.).*/
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%6 == 0 && i%9 == 0 {
			fmt.Println(i)
		}
	}

}
