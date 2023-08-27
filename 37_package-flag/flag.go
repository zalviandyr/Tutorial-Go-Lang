package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your host")
	port := flag.Int("port", 9000, "Put your port, must be number")

	flag.Parse()

	fmt.Println(*host)
	fmt.Println(*port)
}
