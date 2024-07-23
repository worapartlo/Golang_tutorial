package main

import (
	"fmt"
)

func main() {

	zodiac := map[int]string{0: "มะเส็ง",
		1:  "มะเมีย",
		2:  "มะแม",
		3:  "วอก",
		4:  "ระกา",
		5:  "จอ",
		6:  "กุน",
		7:  "ชวด",
		8:  "ฉลู",
		9:  "ขาล",
		10: "เถาะ",
		11: "มะโรง"}

	var year int
	var n_year int

	fmt.Print("ใส่ปี พ.ศ. เกิดของคุณ: ")
	fmt.Scan(&year)
	n_year = year - ((year / 12) * 12)
	//fmt.Println(n_year)
	fmt.Println("ปีนักษัตรของคุณ คือ ", zodiac[n_year], " !")

}
