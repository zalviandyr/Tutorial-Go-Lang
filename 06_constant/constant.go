package main

import "fmt"

func main() {
	const firstName string = "Zukron"
	const lastName = "Alviandy"
	const age = 32

	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		name     string = "Zukron Alviandy"
		nickName        = "alvin"
	)
	fmt.Println(name)
	fmt.Println(nickName)
}
