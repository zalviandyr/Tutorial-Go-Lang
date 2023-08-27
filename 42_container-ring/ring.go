package main

import (
	"container/ring"
	"fmt"
)

func main() {
	data := ring.New(5)

	data.Value = "55"
	data = data.Next()
	data.Value = "3"
	data = data.Next()
	data.Value = 5
	data = data.Next()

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
