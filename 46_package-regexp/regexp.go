package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("a.")

	fmt.Println(re.FindAllString("paranormal", -1))
	fmt.Println(re.FindAllString("paranormal", 1))
	fmt.Println(re.FindAllString("lupa", 1))
}
