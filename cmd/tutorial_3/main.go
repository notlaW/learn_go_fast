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

	var intArray = [3]int{1, 2, 3} // shorthand for creating an array

	fmt.Println(intArray)

	intArrayShorterHand := [3]int{1, 2, 3} // shorter-hand for creating an array
	fmt.Println(intArrayShorterHand)

	intArrayEvenShorterHand := [...]int{1, 2, 3} // even shorter-hand for creating an array with size inference
	fmt.Println(intArrayEvenShorterHand)

	var mySlice []int
	mySlice = append(mySlice, 1)
	fmt.Println(mySlice)

	var myMap = map[string]int{"Adam": 1, "Eve": 2}
	fmt.Println(myMap)

	var age, ok = myMap["Adam"]

	if ok {
		fmt.Printf("Adam's age is %d\n", age)
	} else {
		fmt.Println("Adam's age is not available")
	}

	var fakeAge, okFake = myMap["fakeName"]
	if !okFake {
		fmt.Println(fakeAge)
		fmt.Println("fakeName not found")
	}

	for name := range myMap {
		fmt.Println(name)
	}

}
