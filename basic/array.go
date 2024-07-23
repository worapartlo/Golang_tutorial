package main

import "fmt"

//array
func main() {

	//การใส่ข้อมูล array แบบหลายบรรทัด
	/*var fruit [3] string
	fruit[0] = "apple"
	fruit[1] = "orange"
	fruit[2] = "peach"
	*/

	//การใส่ข้อมูล array แบบบรรทัดเดียว
	fruit := [3]string{"apple", "orange", "peach"}
	amount := [3]int{2, 4, 5}
	fmt.Println(fruit, amount)

}
