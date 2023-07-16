package main

import "fmt"

func main() {
	var persons = [...]string{
		"putra",
		"amel",
		"budi",
	}

	fmt.Println(persons)
	fmt.Println(persons[:2])
	fmt.Println(persons[1:])
	fmt.Println(persons[1:2])
}
