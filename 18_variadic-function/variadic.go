package main

import "fmt"

func main() {
	total := sum(5, 4, 3, 2, 1)
	fmt.Println("Total nya adalah", total)

	label, total2 := sumWithLabel("A", 1, 2, 3)
	fmt.Println("Label", label, "dengan total", total2)
}

func sum(numbers ...int) (total int) {
	total = 0
	for _, number := range numbers {
		total += number
	}

	return
}

func sumWithLabel(label string, numbers ...int) (returnedLabel string, total int) {
	total = sum(numbers...)

	return label, total
}
