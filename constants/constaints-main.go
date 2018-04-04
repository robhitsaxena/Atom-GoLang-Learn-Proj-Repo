package main

import "fmt"

//Declaring Constant
const Name = "Robhit Saxena"

//Multiple declaration of constants
const (
	Pi     = 3.14
	MyName = "Robhit Saxena"
	MyDob  = 19041991
)

func main() {
	const dob = 19041991
	fmt.Printf("%d \n", dob)
	fmt.Printf("%s \n", Name)
	fmt.Printf("%d%s \n", dob, " is age")
}
