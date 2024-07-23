// This Go program defines a simple HTTP server that manages a list of courses, allowing GET and POST
// requests to retrieve and add courses respectively.

package main

import (
	"encoding/json" // ใช้สำหรับการเข้ารหัสและถอดรหัสข้อมูลในรูปแบบ JSON
	"fmt"
	"io"       // ใช้สำหรับการอ่านข้อมูลจาก request body
	"log"      // ใช้สำหรับการบันทึก log
	"net/http" // ใช้สำหรับการสร้าง HTTP server
	"strconv"
	"strings"
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
func findID(ID int) (*Course, int) {
	for i, course := range CourseList { // ลูปผ่านทุก course ใน CourseList
		if course.Id == ID { // ถ้า Id ของ course ตรงกับ ID ที่ต้องการหา
			return &course, i // คืนค่า course และ index ของ course ใน CourseList
		}
	}
	return nil, 0 // ถ้าไม่พบ course ที่มี Id ตรงกัน คืนค่า nil และ 0
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "course/")          // แยกส่วนของ URL เพื่อนำค่า Id ออกมา
	Id, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1]) // แปลงส่วนสุดท้ายของ URL เป็นจำนวนเต็ม
	if err != nil {                                                  // ถ้ามี error ในการแปลงค่า
		log.Print(err)                     // บันทึก error ใน log
		w.WriteHeader(http.StatusNotFound) // ส่งสถานะ 404
		return
	}
	course, listItemIndex := findID(Id) // ค้นหา course จาก Id
	if course == nil {                  // ถ้าไม่พบ course ที่มี Id ตรงกัน
		http.Error(w, fmt.Sprintf("no course with id %d", Id), http.StatusNotFound) // ส่งข้อความ error และสถานะ 404
		return
	}
	switch r.Method {
	case http.MethodGet: // ตรวจสอบว่าคือ GET method หรือไม่
		courseJSON, err := json.Marshal(course) // แปลง course เป็น JSON
		if err != nil {                         // ถ้ามี error ในการแปลงค่า
			http.Error(w, err.Error(), http.StatusInternalServerError) // ส่งข้อความ error และสถานะ 500
			return
		}
		w.Header().Set("Content-type", "application/json") // ตั้งค่า content type เป็น JSON
		w.Write(courseJSON)                                // ส่งข้อมูล JSON ให้กับ client
	case http.MethodPut: // ตรวจสอบว่าคือ PUT method หรือไม่
		var updateCourse Course             // สร้างตัวแปร updateCourse เพื่อเก็บข้อมูล course ที่อัพเดท
		byteBody, err := io.ReadAll(r.Body) // อ่านข้อมูลจาก request body
		if err != nil {                     // ถ้ามี error ในการอ่านข้อมูล
			w.WriteHeader(http.StatusBadRequest) // ส่งสถานะ 400
			return
		}
		err = json.Unmarshal(byteBody, &updateCourse) // แปลงข้อมูล JSON เป็น updateCourse
		if err != nil {                               // ถ้ามี error ในการแปลงค่า
			w.WriteHeader(http.StatusBadRequest) // ส่งสถานะ 400
			return
		}
		if updateCourse.Id != Id { // ถ้า Id ใน updateCourse ไม่ตรงกับ Id ใน URL
			w.WriteHeader(http.StatusBadRequest) // ส่งสถานะ 400
			return
		}
		course = &updateCourse              // อัพเดท course ด้วยข้อมูลใหม่
		CourseList[listItemIndex] = *course // อัพเดท CourseList ด้วย course ที่อัพเดท
		w.WriteHeader(http.StatusOK)        // ส่งสถานะ 200
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed) // ส่งสถานะ 405 ถ้า method ไม่ใช่ GET หรือ PUT
	}
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
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

func enableCorsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization, X-Custom-Header, x-requested-with")
		handler.ServeHTTP(w, r)
	})
}

func main() {
	courseItem := http.HandlerFunc(courseHandler)             // แปลง courseHandler เป็น http.HandlerFunc
	courseList := http.HandlerFunc(coursesHandler)            // แปลง coursesHandler เป็น http.HandlerFunc
	http.Handle("/course/", enableCorsMiddleware(courseItem)) // กำหนด route /course/ ให้ใช้ middlewareHandler กับ courseItem
	http.Handle("/course", enableCorsMiddleware(courseList))  // กำหนด route /course ให้ใช้ middlewareHandler กับ courseList
	http.ListenAndServe(":5000", nil)                         // เริ่มต้น HTTP server ที่พอร์ต 5000
}
