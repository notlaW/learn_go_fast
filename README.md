# Learn Go Fast

Credit to Alex Mux
https://www.youtube.com/watch?v=8uiZC0l4Ajw

# Pre Reqs

## GoLang Overview

- Statically Typed Language
  - Declare variable type or infer.
- Strongly Typed
  - Cannot mix types like you can in nodejs
- Compiled - generally faster than interpreted languages
- Fast compilation time (quicker feedback loop)
- Built in concurrency
- Simplicity
- Powerful standard library

## Packages

A folder that contains a collection of golang files. All go files within a package folder must include the package declaration at the top of the file, and all files in the package must have a matching declaration.

## Modules

A collection of packages is a module. Starting a new golang project is essentially just starting a new module.

## Starting a new project

Initialize the project creating a new module with your choice of project name.

```
$ mkdir learn_go_fast && cd learn_go_fast
$ go mod init learn_go_fast
```

This will create a `go.mod` file that describes the module's properties, including its dependencies on other modules and on versions of Go.

Example `go.mod` file:

````module example.com/mymodule

go 1.14

require (
    example.com/othermodule v1.2.3
    example.com/thismodule v1.2.3
    example.com/thatmodule v1.2.3
)

replace example.com/thatmodule => ../thatmodule
exclude example.com/thismodule v1.3.0```
````

## Main Function

Every go file needs to live within a package. To create an entrypoint function, include `package main` in your main file (typically named `main.go`) which will force the compiler to look for a main function implementation. This main function will be the entrypoint of the compiled executable.

```
package main

func main() {
	println("Hello, World!")
}
```

## Constants, Variables, and Basic Data Types

- Unused variables break compilation

### Number Types

- Cannot mix arithmatic and types

```
package main

import "fmt"

func main() {
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
}
```

## Arrays

Arrays are fixed length, single type, indexable and contiguous in memory. Elements within an array will always be stored together in memory.

## Slices

```
Slices are just wrappers around arrays. 
		- Golang Docs
```


## Things I want to learn in Go 

- API / lambda handler construction
- SQS consumer pattern
- Containerization + build command flows
- Testing
- Productionization of project
	- linting
	- formatting
	- githooks for ^^


Check the tutorial files for more concrete syntax examples. Will try to update the README as I go.


