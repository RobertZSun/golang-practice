package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@go.dev",
			zipCode: 12345,
		},
	}
	// fmt.Printf("%+v", alex)

	alexPointer := &alex
	alexPointer.updateName("Alex1")
	// alex.updateName("Alex2")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (personPointer *person) updateName(newName string) {
	(*personPointer).firstName = newName
}

// func (p person) updateName(newName string) {
// 	p.firstName = newName
// }
