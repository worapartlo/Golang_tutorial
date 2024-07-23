package main

import "fmt"

//map
func main() {

	//การสร้างตัวแปร map ใช้คำสั่ง make(map[datatype]datatype)
	var product = make(map[string]float64)
	fmt.Println(product)

	//การเพิ่มข้อมูลแบบหลายบรรทัด
	product["phone"] = 1500
	product["notebook"] = 250
	product["pen"] = 100
	fmt.Println("list =", product)

	//การลบข้อมูลออก
	delete(product, "phone")
	fmt.Println("list =", product)

	//การเปลี่ยนแปลงค่าที่ map ไว้
	product["pen"] = 200
	fmt.Println("list =", product)

	//การเรียกดูเฉพาะค่าของข้อมูล
	value := product["notebook"]
	fmt.Println("value notebook = ", value)

	//การสร้างตัวแปร map แบบรรทัดเดียว โดยใส่พร้อมข้อมูล
	fruit := map[string]string{"apple": "red", "orange": "orange", "watermelom": "green"}
	fmt.Println(fruit)
}
