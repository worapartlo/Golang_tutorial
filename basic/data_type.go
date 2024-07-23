package main

import "fmt"

//การใช้งาน datatype
func main() {

	//การประกาศตัวแปร รูปแบบที่ 1 var [variable name] [datatype] = [data]
	var name string = "Worapart"
	var age int = 25
	var gpa float32 = 2.27
	var pass bool = true
	const txt string = "I'm a handsome man." //type const เป็นค่าคงที่ไม่สามารถเปลี่ยนค่าได้

	//การแสดงผลข้อมูล รูปแบบที่ 1
	fmt.Println("My name is ", name)
	fmt.Println("My age is ", age)
	fmt.Println("My gpa is ", gpa)
	fmt.Println("My pass is ", pass)
	fmt.Println(txt)
	fmt.Println("**********\n")

	//การประกาศตัวแปร รูปแบบที่ 2 [variable name] := [data] ภาษาจะรู้เองว่า type อะไร
	fullname := "Worapart Loma-in"
	height := 175
	weight := 70.05
	fat := false

	//การแสดงผลข้อมูล รูปแบบที่ 2 การใช้ Printf
	//Printf จะไม่ขึ้นบรรทัดใหม่ต้อง \n
	//%v แสดงค่าในตัวแปร %T แสดง type ข้อมูล
	fmt.Printf("I am %v\n", fullname)
	fmt.Printf("I tall %v cm ", height)
	fmt.Printf("My weight is %v kg\n", weight)
	fmt.Printf("I am fat: %v\n", fat)
	fmt.Println("**********\n")

	fmt.Printf("data type %v : %T\n", fullname, fullname)
	fmt.Printf("data type %v : %T\n", age, age)
	fmt.Printf("data type %v : %T\n", gpa, gpa)
	fmt.Printf("data type %v : %T\n", pass, pass)
}
