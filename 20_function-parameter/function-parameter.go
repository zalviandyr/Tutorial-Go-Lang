package main

import "fmt"

type Filter func(string) string

func main() {
	sayHelloWithFilter("Nae", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

	// Anonymous function
	sayHelloWithFilter("Goblok", func(name string) string {
		if name == "Goblok" {
			return "Lu yang goblok"
		} else {
			return name
		}
	})
}

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)

	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
