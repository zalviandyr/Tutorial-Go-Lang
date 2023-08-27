package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Zukron Alviandy"
	nameCaps := strings.ToUpper(name)
	nameSplit := strings.Split(name, " ")

	fmt.Println(name)
	fmt.Println(nameCaps)
	fmt.Println(nameSplit)
}
