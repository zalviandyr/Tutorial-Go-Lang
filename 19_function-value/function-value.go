package main

import "fmt"

func main() {
	goodBye := getGoodBye
	fmt.Println(goodBye("Alvin"))

	var goodBye2 func(string) string = getGoodBye
	fmt.Println(goodBye2("Zukron"))

	var goodBye3 = getGoodBye
	fmt.Println(goodBye3("Alviandy"))
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}
