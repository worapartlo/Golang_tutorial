package main

import "fmt"

func tempCtoF(val float64) {
	var result float64
	result = (val * (9 / 5)) + 32
	fmt.Println("Temperature in Fahrenheit: ", result)
}

func tempFtoC(val float64) {
	var result float64
	result = ((val - 32) * 5) / 9
	fmt.Println("Temperature in Celsius: ", result)
}

func main() {
	fmt.Println("*****Simple Temperature converter*****")

	var temp string
	var tempcon float64

	fmt.Print("Select Temperature to convert C(c to f) or F(f to c):")
	fmt.Scan(&temp)
	if temp == "c" || temp == "C" {
		fmt.Print("Temperature in Celsius: ")
		fmt.Scan(&tempcon)
		tempCtoF(tempcon)
	} else if temp == "f" || temp == "F" {
		fmt.Print("Temperature in Fahrenheit: ")
		fmt.Scan(&tempcon)
		tempFtoC(tempcon)
	} else {
		panic("error please type only C or F...")
	}

}
