package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID    int
	Name  string
	Email string
}

func main() {
	data, _ := json.Marshal(&employee{001, "John Smith", "john_m@gmail.com"})
	fmt.Println(string(data))
}
