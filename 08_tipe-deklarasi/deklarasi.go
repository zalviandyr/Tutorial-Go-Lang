package main

import "fmt"

func main() {
	type NoKTP string

	var noKtp NoKTP = "133313131"
	fmt.Println(noKtp)
}
