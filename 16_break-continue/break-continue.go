package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Index ke ", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Index ke ", i)
	}
}
