package main

import "fmt"

func main() {
	// declare and init with explicitly type
	var coffeeName string = "Espresso"

	// type inferred
	var size = "Small"
	// short declaration and init
	var price = 2.50

	fmt.Println("Small Espresso price is $2.50")
	fmt.Println(size, coffeeName, "price is $", price)
	fmt.Printf("%s %s price is $%.2f\n", size, coffeeName, price)

}
