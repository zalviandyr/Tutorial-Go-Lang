package main

import "fmt"

func main() {
	result := random()
	fmt.Println(result.(string))

	//resultInt := random()
	//fmt.Println(resultInt.(int))

	switch result.(type) {
	case string:
		fmt.Println("string", result)
	case int:
		fmt.Println("int", result)
	default:
		fmt.Println("unknown")
	}
}

func random() any {
	return "Ok"
}
