package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type engine interface {
	milesLeft() uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("You can't make it!")
	}
}

type electricEngine struct {
	mpkhw uint8
	kwh   uint8
}

// func hasEnoughGas(e gasEngine, miles int) {
// 	if miles <= e.milesLeft(e) {
// 		fmt.Println("You have enough gas")
// 	} else {
// 		fmt.println("You don't have enough gas")
// 	}
// }

func main() {
	fmt.Println("Tutorial 5: Structs and Interfaces")

	var myEngine gasEngine = gasEngine{mpg: 30.0, gallons: 10.0}
	fmt.Println(myEngine)
	fmt.Println("Miles left in the tank:", myEngine.milesLeft())
	miles := 10

	fmt.Printf("For %d miles, can I make it?\n", miles)
	canMakeIt(myEngine, uint8(miles))

}
