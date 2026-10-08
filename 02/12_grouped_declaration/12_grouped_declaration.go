package main

import "fmt"

func main() {
	var coffeeType string = "Latte"
	var quantity int = 3
	var unitPrice float64 = 5.50

	var (
		customerName string = "Dima"
		tableNumber  int    = 3
		isReadyToPay bool   = false
	)

	fmt.Printf("Ordered %d %s priced $%.2f each\n", quantity, coffeeType, unitPrice)
	fmt.Printf("Customer %s at table %d is ready to pay: %t\n", customerName, tableNumber, isReadyToPay)

	//no compilation error on unused const
	const (
		SizeSmall  = 'S'
		SizeMedium = 'M'
		SizeLarge  = 'L'
	)

}
