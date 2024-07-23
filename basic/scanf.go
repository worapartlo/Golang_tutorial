package main

import "fmt"

//รับข้อมูลจาก keyboard
func main() {

	//สร้าง variable เพื่อเก็บข้อมูล
	var name string
	var num float32

	//การถามและเก็บข้อมูลจากผู้ใช้
	fmt.Print("Welcome what your name: ")
	fmt.Scan(&name)

	fmt.Print("What is your height in cm: ")
	fmt.Scan(&num)

	//แสดงผลข้อมูลที่เก็บมา
	fmt.Println("\nHello mr.", name)
	fmt.Println("Your height is", num/100)
}
