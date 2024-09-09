package main

import "fmt"

func main() {

	fmt.Println("Tutorial 1: Variables")


	//Numbers
	var intNumber int
	fmt.Println(intNumber)
	var intNumber32 int32
	fmt.Println(intNumber32)
	var intNumber64 int64
	fmt.Println(intNumber64)
	var floatNumber float32
	fmt.Println(floatNumber)
	var float64Number float64
	fmt.Println(float64Number)
	var unsigned uint64
	fmt.Println(unsigned)

	//Strings
	var stringText string = "Hello, World!"
	fmt.Println(stringText)
	var concatStringText string = "Hello" + " " + "World!"
	fmt.Println(concatStringText)
	fmt.Printf("My robot says: %s", concatStringText)
	var multiLineString string = `Hello


	world!`
	
	fmt.Printf(multiLineString)

	//Booleans
	var isTrue bool = true
	fmt.Println(isTrue)

	//Shorthand
	shortHandString := "Hello, World!"
	fmt.Println(shortHandString)
	myvar1, myvar2 := 1, 2
	fmt.Println(myvar1, myvar2)

	//Constants
	const myConst = "Hello, World!"

	// ^^ Does not break compilation when not used.
	//Arrays
	var myArray [3]int
	fmt.Println(myArray)
}