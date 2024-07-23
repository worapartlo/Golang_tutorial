package main

import "fmt"

//pointer
func zeroval(ival int) {
	ival = 0
}

//function pointer ต้องระบุ * เป็นตัวชี้ pointer
func zeropoint(ipoint *int) {
	*ipoint = 0
}
func main() {
	i := 1
	fmt.Println("i = ", i)

	//ค่าที่ส่งไปจะไม่เปลี่ยนแปลงกลับมา
	zeroval(i)
	fmt.Println("i from function zero =", i)

	//การส่งค่าไปต้องใส่ & ไปด้วย
	zeropoint(&i)
	fmt.Println("i value from function zero =", i)
	fmt.Println("i address from function zero =", &i)
}
