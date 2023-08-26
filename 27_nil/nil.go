package main

import "fmt"

func main() {
	data := NewMap("Alvin")
	fmt.Println(data)

	data2 := NewMap("")
	fmt.Println(data2)
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"Name": name,
		}
	}
}
