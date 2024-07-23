package main

import "fmt"

//ตัวดำเนินการทางคณิตศาสตร์
func main() {

	//ประกาศตัวแปร
	x := 10
	y := 8

	//การดำเนินการ + - * / %
	fmt.Println("ผลบวก = ", x+y)
	fmt.Println("ผลลบ = ", x-y)
	fmt.Println("ผลคูณ = ", x*y)
	fmt.Println("ผลหาร = ", x/y)
	fmt.Println("เศษ = ", x%y)

	//ประกาศตัวแปรที่มี datatype เหมือนกัน
	var i, j = 2, 3 //หรือ i, j := 2, 3

	//การดำเนินการการเปรียบเทียบ
	fmt.Println("เท่ากัน:", i == j)
	fmt.Println("ไม่เท่ากัน:", i != j)
	fmt.Println("น้อยกว่า:", i < j)
	fmt.Println("มากกว่า:", i > j)
	fmt.Println("น้อยกว่าเท่ากับ:", i <= j)
	fmt.Println("มากกว่าเท่ากับ:", i >= j)
}
