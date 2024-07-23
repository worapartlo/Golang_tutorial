package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID    int
	Name  string
	Email string
}

func main() {

	e := employee{}
	err := json.Unmarshal([]byte(`{"ID":1,"Name":"John Smith","Email":"john_m@gmail.com"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)
	fmt.Println(e.Name)
}
