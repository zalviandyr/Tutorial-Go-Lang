package main

import "fmt"

func main() {
	sayHello("Alvin")

	a := 5
	b := 7
	result := add(a, b)
	fmt.Printf("Hasil %d + %d = %d\n", a, b, result)

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	fmt.Println(getFullName())

	umur, nama := getData()
	fmt.Printf("Umur %s adalah %d tahun\n", nama, umur)
}

func sayHello(name string) {
	fmt.Println("Hai", name)
}

func add(a int, b int) int {
	return a + b
}

func getFullName() (string, string) {
	return "Zukron", "Alviandy"
}

func getData() (age int, name string) {
	age = 22
	name = "Alvin"

	return
}
