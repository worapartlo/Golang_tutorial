package main

import "fmt"

func main() {
	i := 0
	for i <= 10 {
		fmt.Println("i =", i)
		i += 1
	}

	count := 0
	for {
		fmt.Println("Hello, world!")
		count += 1
		if count == 20 {
			break
		}
	}
}
