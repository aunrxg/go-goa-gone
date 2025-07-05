package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Print("Enter the Number: ")
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {

		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}
}
