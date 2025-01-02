package main

import "fmt"

func main() {

	userNames := []string{}

	userNames = append(userNames, "Sally")
	userNames = append(userNames, "John")
	userNames = append(userNames, "Jane")

	fmt.Println(userNames)

}
