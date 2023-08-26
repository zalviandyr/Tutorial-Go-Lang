package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func main() {
	alvin := Person{
		Name: "Alvin",
	}

	sayHello(alvin)

	kucing := Animal{
		Name: "Kucing",
	}

	sayHello(kucing)
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func sayHello(hasName HasName) {
	fmt.Println("Hello,", hasName.GetName())
}
