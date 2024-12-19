package main

import "fmt"

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter Revenue, Expenses, Tax Rate: ")
	fmt.Scan(&revenue, &expenses, &taxRate)

	ebt := revenue - expenses
	eat := ebt * (1 - taxRate/100)
	ratio := ebt / eat
	tax := ebt * (taxRate / 100)

	fmt.Println("Hello World: ", ebt, eat, ratio, tax)
}
