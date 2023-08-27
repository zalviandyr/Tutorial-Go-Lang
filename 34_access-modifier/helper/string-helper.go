package helper

import "fmt"

// private
func sayHello(name string) {
	fmt.Println("Hello", name)
}

// public
func SayHello(name string) {
	fmt.Println("Hello", name)
}
