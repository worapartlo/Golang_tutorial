// This Go program defines a simple HTTP server that manages a list of courses, allowing GET and POST
// requests to retrieve and add courses respectively.

package main

import (
	"encoding/json" // ใช้สำหรับการเข้ารหัสและถอดรหัสข้อมูลในรูปแบบ JSON
	"io"            // ใช้สำหรับการอ่านข้อมูลจาก request body
	"log"           // ใช้สำหรับการบันทึก log
	"net/http"      // ใช้สำหรับการสร้าง HTTP server
)

type Course struct {
	Id         int    `json: "id"`         // ฟิลด์ Id ของ course
	Name       string `json: "name"`       // ฟิลด์ Name ของ course
	Price      int    `json: "price"`      // ฟิลด์ Price ของ course
	Instructor string `json: "instructor"` // ฟิลด์ Instructor ของ course
}

var CourseList []Course // สร้างตัวแปร CourseList เป็น slice ของ Course

func init() {
	CourseJSON := `[
		{"id":101, "name":"Python", "price":1250, "instructor":"Worapart"},
		{"id":102, "name":"HTML", "price":1000, "instructor":"Worapart"},
		{"id":103, "name":"C++", "price":950, "instructor":"Worapart"}
	]` // กำหนดค่าเริ่มต้นให้กับ CourseList ในรูปแบบ JSON
	err := json.Unmarshal([]byte(CourseJSON), &CourseList) // แปลง JSON เป็น slice ของ Course
	if err != nil {
		log.Fatal(err) // ถ้ามี error ให้แสดง error แล้วหยุดการทำงาน
	}
}

func getNextID() int {
	highestID := -1                     // เริ่มต้น highestID ด้วยค่าต่ำสุด
	for _, course := range CourseList { // ลูปผ่านทุก course ใน CourseList
		if highestID < course.Id { // ถ้า Id ของ course ปัจจุบันมากกว่า highestID
			highestID = course.Id // อัพเดทค่า highestID
		}
	}
	return highestID + 1 // คืนค่า highestID ที่เพิ่มขึ้น 1
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	courseJSON, err := json.Marshal(CourseList) // แปลง CourseList เป็น JSON
	switch r.Method {
	case http.MethodGet: // ตรวจสอบว่าคือ GET method หรือไม่
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError) // ถ้ามี error ให้ส่งสถานะ 500
			return
		}
		w.Header().Set("Content-Type", "application/json") // ตั้งค่า content type เป็น JSON
		w.Write(courseJSON)                                // ส่งข้อมูล JSON ให้กับ client

	case http.MethodPost: // ตรวจสอบว่าคือ POST method หรือไม่
		var newCourse Course                // สร้างตัวแปร newCourse เพื่อเก็บข้อมูล course ใหม่
		Bodybyte, err := io.ReadAll(r.Body) // อ่านข้อมูลจาก request body
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) // ถ้ามี error ให้ส่งสถานะ 400
			return
		}
		err = json.Unmarshal(Bodybyte, &newCourse) // แปลงข้อมูล JSON เป็น newCourse
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) // ถ้ามี error ให้ส่งสถานะ 400
			return
		}
		if newCourse.Id != 0 {
			w.WriteHeader(http.StatusBadRequest) // ถ้า Id ไม่ใช่ 0 ให้ส่งสถานะ 400
			return
		}
		newCourse.Id = getNextID()                 // กำหนด Id ใหม่ให้กับ newCourse
		CourseList = append(CourseList, newCourse) // เพิ่ม newCourse ลงใน CourseList
		w.WriteHeader(http.StatusCreated)          // ส่งสถานะ 201
		return
	}
}

func main() {
	http.HandleFunc("/course", courseHandler) // กำหนด route /course ให้กับ courseHandler
	http.ListenAndServe(":5000", nil)         // เริ่มต้น HTTP server ที่พอร์ต 5000
}
