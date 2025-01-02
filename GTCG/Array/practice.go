package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {

	// task 1
	myHobbies := []string{"walking", "swimming", "rowing"}
	fmt.Println(myHobbies)

	// task 2
	fmt.Println(myHobbies[0])
	fmt.Println(myHobbies[1:3])

	// task 3
	myHobbiesSlice := myHobbies[0:2]
	myHobbiesSlice2 := []string{myHobbies[0], myHobbies[1]}
	myHobbiesSlice3 := append([]string{}, myHobbies[0], myHobbies[1])

	fmt.Println(myHobbiesSlice, myHobbiesSlice2, myHobbiesSlice3)

	// task 4
	myHobbiesSlice = myHobbiesSlice[1:3]

	fmt.Println("task 4: ", myHobbiesSlice)

	// task 5
	myGoals := []string{"Learn Golang", "Learn Python"}
	fmt.Println("task 5: ", myGoals)

	// task 6
	myGoals[1] = "Learn Java"
	myGoals = append(myGoals, "Learn C++")

	fmt.Println("task 6: ", myGoals)

	// task 7
	// myProducts := []Product{
	// 	Product{id: "1", title: "Product 1", price: 1.99},
	// 	Product{id: "2", title: "Product 2", price: 2.99},
	// }
	myProducts := []Product{
		Product{"first", "first-product", 1.99},
		Product{"second", "second-product", 1.99},
	}

	myProducts = append(myProducts, Product{id: "3", title: "Product 3", price: 3.99})

	fmt.Println("task 7: ", myProducts)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
