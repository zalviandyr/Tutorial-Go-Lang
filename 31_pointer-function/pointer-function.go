package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	person1 := Person{"Pak Le"}
	person2 := &person1

	changeName(person2)

	fmt.Println(person1)
	fmt.Println(person2)
}

func changeName(person *Person) {
	person.Name = "Buk Le"
}
