package main

import "fmt"

const (
	water       = 200
	milk        = 50
	coffeeBeans = 15
)

func makeCoffee(cups uint) {
	fmt.Printf("For %d cups of coffee you will need:\n", cups)
	fmt.Printf("%d ml of water\n", water*cups)
	fmt.Printf("%d ml of milk\n", milk*cups)
	fmt.Printf("%d g of coffee beans\n", coffeeBeans*cups)
}

func main() {
	fmt.Println("Write how many cups of coffee you will need:")

	var cups uint
	fmt.Scan(&cups)

	makeCoffee(cups)
}
