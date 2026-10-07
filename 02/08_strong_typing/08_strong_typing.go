package main

import "fmt"

func main() {
	price := 4.50
	quantity := 15

	// total := price * quantity // mismatched
	total := price * float64(quantity)
	fmt.Printf("Total income: %.2f", total)

	total2 := int(price) * quantity
	fmt.Printf("Total income: %s", total2)
}
