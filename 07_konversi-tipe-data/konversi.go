package main

import "fmt"

func main() {
	var nilai32 int32 = 1331
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64)

	fmt.Println("Nilai 16 =", nilai16)

	fmt.Println(string("Zukron"[0]))
}
