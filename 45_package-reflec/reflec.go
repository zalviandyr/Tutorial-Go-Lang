package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person1 := Person{"Bahlul", 23}
	person2 := Person{"Bahlul", 23}
	person3 := Person{"Ucup", 23}

	fmt.Println("person1 is equal person2:", reflect.DeepEqual(person1, person2))
	fmt.Println("person1 is not equal person3:", reflect.DeepEqual(person1, person3))
}
