package main

import (
	"errors"
	"fmt"
)

func printMe(tutorialName string) {
	fmt.Printf("Tutorial 2: %s\n", tutorialName)
}

func intDivision(numerator, demoninator int) int {
	return numerator / demoninator
}

func intDivisionWithRemainder(numerator, demoninator int) (int, int) {
	result := numerator / demoninator
	remainer := numerator % demoninator
	return result, remainer
}

func intDivisionWithError(numerator, demoninator int) (int, int, error) {
	var err error // inits as nil (a lot of things init as nil)
	if demoninator == 0 {
		//because we have to adhere to the return type, we have to return 0, 0 and then the error.
		//The error must be included in the signature above.
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}

	result := numerator / demoninator
	remainder := numerator % demoninator
	return result, remainder, err
}

func main() {
	fmt.Println("Tutorial 3: Functions and control flow")

	var result, remainder, err = intDivisionWithError(10, 0)

	if err != nil {
		fmt.Println("An error occurred: ")
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("Result: %d\n", result)
	} else {
		fmt.Printf("Result: %d, Remainder: %d\n", result, remainder)
	}

	if result == 0 && remainder == 0 {
		fmt.Println("Both result and remainder are 0")
	}

	// Switch statement
	switch {
	case err != nil:
		fmt.Println("An error occurred: ")
		fmt.Printf(err.Error())
	case result == 0 && remainder == 0:
		fmt.Println("Both result and remainder are 0")
	case result == 0:
		fmt.Println("Result is 0")
	case remainder == 0:
		fmt.Println("Remainder is 0")
	default:
		fmt.Println("Result and remainder are not 0")
	}
}
