package main

import "fmt"

func main() {
	var targetSlice []int

	for i := 1; i <= 10; i++ {
		targetSlice = append(targetSlice, i)
	}

	for value := range targetSlice {
		if value%2 == 0 {
			fmt.Println(value, "is even")
		} else {
			fmt.Println(value, "is odd")
		}
	}

	slice := []int{10, 20, 30}

	// Index and value
	for i, v := range slice {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	for v := range slice {
		fmt.Println("Value:", v)
	}
}
