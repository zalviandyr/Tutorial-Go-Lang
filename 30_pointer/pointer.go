package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Bungo", "Jambi", "Indonesia"}
	address2 := &address1

	address2.City = "Kerinci"

	*address2 = Address{"Yogyakarta", "DIY Yogyakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	address4 := new(Address)
	address4.City = "Jakarta"

	address5 := address4

	fmt.Println(address4)
	fmt.Println(address5)
}
