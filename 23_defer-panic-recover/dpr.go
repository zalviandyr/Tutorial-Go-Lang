package main

import "fmt"

func main() {
	runApplication(9000)
}

func runApplication(port int) {
	// always executed when error
	defer logging()

	// to catch panic
	defer endApp()

	if port == 9000 {
		panic("Port has been used")
	}

	fmt.Println("Run Application")
}

func logging() {
	fmt.Println("Selesai memanggil function")
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Terjadi error:", message)
	}
}
