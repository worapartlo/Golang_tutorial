package main

import "fmt"

//sclices
func main() {

	//slices จะเก็บข้อมูลโดยไม่ต้องระบุขอบเขต
	color := []string{"green", "red", "yellow"}
	fmt.Println(color)

	//การเก็บข้อมูลสามารถเพิ่มได้เรื่อยๆ
	color = append(color, "blue", "pink", "light blue")
	fmt.Println(color)

	//การเลือกแสดงผลข้อมูล
	new_color := color[2:5]
	fmt.Println(new_color)

	new_color = color[0:2]
	fmt.Println(new_color)
}
