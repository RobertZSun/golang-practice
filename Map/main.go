package main

import "fmt"

func main() {

	// var myMao map[string]string
	// myMap := make(map[string]string)
	// myMap := map[string]string{
	// 	"foo":  "bar",
	// 	"foo1": "bar1",
	// }

	// fmt.Println("Welcome to Maps")
	// printMap(myMap)

	m := map[string]string{
		"dog": "bark",
		"cat": "purr",
	}

	for key, value := range m {
		fmt.Println(value)
	}
}

func printMap(c map[string]string) {

	for key, value := range c {
		fmt.Println(key, value)
	}
}
