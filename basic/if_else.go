package main

import "fmt"

//if else condition
func main() {
	var gpa int
	fmt.Println("Grade calculator")
	fmt.Print("Enter your gpa: ")
	fmt.Scan(&gpa)
	fmt.Println("Your gpa =", gpa)

	//if else จะคล้ายภาษา python ไม่ต้องมีวงเล็บที่เงื่อนไข
	if gpa >= 80 {
		fmt.Println("Your got A grade!")
	} else if gpa >= 70 {
		fmt.Println("Your got B grade!")
	} else if gpa >= 60 {
		fmt.Println("Your got C grade!")
	} else if gpa >= 50 {
		fmt.Println("Your got D grade!")
	} else {
		fmt.Println("Your got F grade!")
	}
}
