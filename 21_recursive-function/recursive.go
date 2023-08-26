package main

import "fmt"

func main() {
	a := 3
	fmt.Println("Factorial", a, "adalah", factorialLoop(a))

	fmt.Println("Factorial", a, "adalah", factorialRecursive(a))
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialLoop(value-1)
	}
}
