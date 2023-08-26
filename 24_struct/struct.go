package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var customer Customer
	customer.Name = "Alvin"
	customer.Address = "Jogga"
	customer.Age = 22

	customer.sayHello()

	customer2 := Customer{
		Name:    "Zukron",
		Address: "Jambi",
		Age:     30,
	}

	customer2.sayHello()

	customer3 := Customer{"Rahmadhan", "Bungo", 22}

	customer3.sayHello()
}

func (customer Customer) sayHello() {
	fmt.Println(customer.Name, "yang berumur", customer.Age, "tinggal di", customer.Address)
}
