package main

import "fmt"

func main() {
	greeting := func() /* It is an anonymous function because it has no name. */ {
		fmt.Println("Hello World")
	}
	greeting()
	fmt.Printf("%T\n", greeting)

	// When we want to have a function inside a function then we use func expression
	// When we assign a function to a variable then it's called func expression.

	greet := returnfunc()
	fmt.Println(greet())
  fmt.Printf("%T \n",greet)
}

func returnfunc() func() string {
	return func() string {
		return "Hello"
	}
}
