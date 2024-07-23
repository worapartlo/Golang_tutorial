package main

import "fmt"

func main() {

	var val1, val2, result int

	fmt.Println("Simple calculator")
	fmt.Print("value1: ")
	fmt.Scan(&val1)
	fmt.Print("value2: ")
	fmt.Scan(&val2)

	var operator string

	fmt.Print("operation ( + - * /): ")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		result = val1 + val2
	case "-":
		result = val1 - val2
	case "*":
		result = val1 * val2
	case "/":
		result = val1 / val2
	}
	fmt.Println("result = ", result)

}
