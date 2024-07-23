package main

import "fmt"

//functions
//ฟังก์ชั่นแสดงผลคำสั่ง
func hello() {
	fmt.Println("Hello everyone!")
}

//ฟังก์ชั่นการบวกเลขและแสดงผลลัพธ์
func plus(val int, val2 int) {
	result := val + val2
	fmt.Println(result)
}

//ฟังก์ชั่นการคูณเลขและส่งค่ากลับไป ฟังก์ชั่นหลัก
func multiply(val int, val2 int) int {
	return val * val2
}

func main() {

	hello() //การเรียนใช้ฟังก์ชั่น

	plus(2, 3) //การส่งค่าไปฟังก์ชั่นเพื่อทำการบวกเลข

	result := multiply(3, 4) //การส่งค่าไปฟังก์ชั่นคูณเลขและส่งค่ากลับมาแสดงผลที่ฟังก์ชั่นหลัก
	fmt.Println("result =", result)
}
