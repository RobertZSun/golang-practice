package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type card struct {
	suit  string
	value int
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
