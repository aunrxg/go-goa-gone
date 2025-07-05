package main

import (
	"fmt"
)

func Calc(num1, num2 int, sign string) (int, error) {

	switch sign {
	case "+":
		return num1+num2, nil
	case "-":
		return num1-num2, nil
	case "*":
		return num1*num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("can not divide by zero")
		}
		return num1/num2,  nil
	default:
		return 0, fmt.Errorf("invalid operation")
	}
}

func main() {
	var num1 int
	var num2 int
	var sign string

	fmt.Print("Enter the First number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the Opearation (+,-,*,/): ")
	fmt.Scan(&sign)

	fmt.Print("Enter the Second number: ")
	fmt.Scan(&num2)


	result, err := Calc(num1, num2, sign)

	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Println("Result: ", result)
}
