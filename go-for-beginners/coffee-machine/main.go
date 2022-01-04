package main

import (
	"fmt"
)

type CoffeeMachine struct {
	water, milk, coffeeBeans, cups, money uint
}

type Coffee struct {
	water, milk, coffeeBeans, cost uint
}

// recipes
var espresso = Coffee{water: 250, milk: 0, coffeeBeans: 16, cost: 4}
var latte = Coffee{water: 350, milk: 75, coffeeBeans: 20, cost: 7}
var cappuccino = Coffee{water: 200, milk: 100, coffeeBeans: 12, cost: 6}
var hotWater = Coffee{water: 250, cost: 1}

func (cm *CoffeeMachine) actionLoop() {
	for cm.doAction() {
	}
}

func (cm *CoffeeMachine) doAction() bool {
	fmt.Println("Write action (buy, fill, take, remaining, exit)")

	var action string
	fmt.Scan(&action)
	fmt.Println()

	switch action {
	case "buy":
		cm.doBuy()
	case "fill":
		cm.doFill()
	case "take":
		cm.doTake()
	case "exit":
		return false
	case "remaining":
		cm.printState()
	default:
		fmt.Println("Wrong action. Available actions: " +
			"buy, fill, take, remaining, exit")
	}

	fmt.Println()

	return true
}

func (cm *CoffeeMachine) doFill() {
	var add uint

	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&add)
	cm.water += add
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&add)
	cm.milk += add
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&add)
	cm.coffeeBeans += add
	fmt.Println("Write how many disposable cups you want to add:")
	fmt.Scan(&add)
	cm.cups += add
}

func (cm *CoffeeMachine) doTake() {
	fmt.Printf("I gave you $%d\n", cm.money)
	cm.money = 0
}

func (cm *CoffeeMachine) doBuy() {
	fmt.Println("What do you want to buy? " +
		"1 - espresso, 2 - latte, 3 - cappuccino, " +
		"4 - hot water, back - to main menu")

	var choice string
	var recipe *Coffee
	fmt.Scan(&choice)

	switch choice {
	case "1":
		recipe = &espresso
	case "2":
		recipe = &latte
	case "3":
		recipe = &cappuccino
	case "4":
		recipe = &hotWater
	case "back":
		return
	default:
		fmt.Println("Wrong type of coffee!")
		return
	}

	cm.buyCoffee(recipe)
}

func infoResourceShortage(resource string) {
	fmt.Printf("Sorry, not enough %s!\n", resource)
}

func (cm *CoffeeMachine) buyCoffee(coffee *Coffee) {
	if cm.water < coffee.water {
		infoResourceShortage("water")
		return
	}

	if cm.milk < coffee.milk {
		infoResourceShortage("milk")
		return
	}

	if cm.coffeeBeans < coffee.coffeeBeans {
		infoResourceShortage("coffee beans")
		return
	}

	if cm.cups == 0 {
		infoResourceShortage("cups")
		return
	}

	fmt.Println("I have enough resources, making you a coffee!")

	cm.water -= coffee.water
	cm.milk -= coffee.milk
	cm.coffeeBeans -= coffee.coffeeBeans
	cm.money += coffee.cost
	cm.cups--
}

func (cm *CoffeeMachine) printState() {
	fmt.Printf("The coffee machine has:\n"+
		"%d of water\n"+
		"%d of milk\n"+
		"%d of coffee beans\n"+
		"%d of disposable cups\n"+
		"%d of money\n",
		cm.water,
		cm.milk,
		cm.coffeeBeans,
		cm.cups,
		cm.money)
}

func main() {
	// initial state
	cm := CoffeeMachine{
		water:       400,
		milk:        540,
		coffeeBeans: 120,
		cups:        9,
		money:       550,
	}

	cm.actionLoop()
}
