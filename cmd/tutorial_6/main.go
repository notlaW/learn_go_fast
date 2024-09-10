package main

import "fmt"

func main() {
	println("Tutorial 6: Pointers")
	// pointers are a special type of variable that stores the memory address of another variable
	var p *int32 // p is a pointer to an int

	fmt.Println(p) // nil pointer at this time

	// nil meaning it doesn't point to anything
	// Another way to say this pointer does not have an address assigned to it, in which to store an int value
	// Pointers take up 32 or 64 bits depending on your os

	p = new(int32) // new is a built-in function that allocates an address to the pointer.
	// It gives us back a free memory location which is 32 bits wide which we can store an int value.

	fmt.Printf("Pointer address %v\n", p) // 0x1400000xxxx
	fmt.Printf("Pointer value %v\n", *p)  // 0

	// Notice that the value of the pointer is 0, this is because we haven't assigned a value to the memory location so it
	// defaults to the zero value of the type it points to.

	var x = 0x140000b2020
	fmt.Println(x)
	fmt.Println(&x) // memory address of x

	fmt.Println(*p) // Get's the value of the memory location that p points to, also called dereferencing

	*p = 10 // Assigning a value to the memory location that p points to

	fmt.Println(*p) // 10

	// Pointers are useful when you want to pass a large data structure to a function without copying it
	// This is because pointers are only 32 or 64 bits in size, so they are cheap to pass around

	// Pointers are also useful when you want to modify a variable in a function and have that change reflected in the caller
	// This is because the pointer points to the same memory location in both the caller and the function

	// Pointers are also useful when you want to return multiple values from a function
	// This is because you can pass a pointer to a variable to a function and have that function modify the value of the variable

	// Pointers are also useful when you want to create a data structure that references itself
	// This is because you can create a pointer to the same type of data structure

}
