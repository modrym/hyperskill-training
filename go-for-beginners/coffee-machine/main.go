package main

import (
	"fmt"
)

const (
	water       = 200
	milk        = 50
	coffeeBeans = 15
)

func min(nums ...uint) uint {
	var minNum = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < minNum {
			minNum = nums[i]
		}
	}

	return minNum
}

func getAvailableIngredients() (uint, uint, uint) {
	var availableWater, availableMilk, availableCoffeeBeans uint

	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scan(&availableWater)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scan(&availableMilk)
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scan(&availableCoffeeBeans)

	return availableWater, availableMilk, availableCoffeeBeans
}

func makeCoffee(cups uint, availableWater uint, availableMilk uint, availableCoffeeBeans uint) {
	availableCups := min(availableWater/water, availableMilk/milk, availableCoffeeBeans/coffeeBeans)

	switch {
	case availableCups == cups:
		fmt.Println("Yes, I can make that amount of coffee")
	case availableCups > cups:
		fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that)\n", availableCups-cups)
	default:
		fmt.Printf("No, I can make only %d cups of coffee\n", availableCups)
	}
}

func specifyCups() uint {
	fmt.Println("Write how many cups of coffee you will need:")

	var cups uint
	fmt.Scan(&cups)

	return cups
}

func main() {
	availableWater, availableMilk, availableCoffeeBeans := getAvailableIngredients()
	cups := specifyCups()
	makeCoffee(cups, availableWater, availableMilk, availableCoffeeBeans)
}
