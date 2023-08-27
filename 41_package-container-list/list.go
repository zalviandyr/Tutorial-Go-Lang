package main

import (
	"container/list"
	"fmt"
)

func main() {
	persons := list.New()
	persons.PushBack("Tree")

	for person := persons.Front(); person != nil; person = person.Next() {
		fmt.Println(person.Value)
	}
}
