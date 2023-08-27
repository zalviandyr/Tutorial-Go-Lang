package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	person1 := Person{"Cika"}
	person2 := Person{"Agustin"}

	person1.Madam()
	person2.Madam()

	fmt.Println(person1)
	fmt.Println(person2)
}

func (person *Person) Madam() {
	person.Name = "Mrs. " + person.Name
}
