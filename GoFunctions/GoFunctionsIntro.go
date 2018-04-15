package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	greet("Robhit", "Saxena")
}

/*
func <receivers> <name> (<parameters>) <return> { <code> }
*/

func greet(name string, surname string) {
	fmt.Println(name, surname)

}
