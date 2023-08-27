package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
}

type PersonSlice []Person

func main() {
	persons := PersonSlice{{"Alvin"}, {"Buya"}, {"Zera"}, {"Cika"}}

	fmt.Println("Length:", persons.Len())
	fmt.Println("Less:", persons.Less(1, 0))

	persons.Swap(0, 1)
	fmt.Println(persons)

	sort.Sort(persons)
	fmt.Println("After sort", persons)
}

func (personSlice PersonSlice) Len() int {
	return len(personSlice)
}

func (personSlice PersonSlice) Less(i int, j int) bool {
	return personSlice[i].Name < personSlice[j].Name
}

func (personSlice PersonSlice) Swap(i int, j int) {
	personSlice[i], personSlice[j] = personSlice[j], personSlice[i]
}
