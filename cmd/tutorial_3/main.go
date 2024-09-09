package main

import "fmt"

func main() {
	fmt.Println("Tutorial 3: Arrays, slices, maps and loops")
	var myArray [3]int

	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	fmt.Println(myArray)
	fmt.Println(myArray[1:2])
	fmt.Println(&myArray[1])

}
