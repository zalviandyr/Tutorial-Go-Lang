package main

import (
	"fmt"
	"strconv"
)

func main() {
	number, _ := strconv.ParseInt("64", 16, 32)

	fmt.Println(number)
}
