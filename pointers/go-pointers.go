package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	var b *int = &a // b has the memory address to the address where int is stored
	fmt.Println(b)
	fmt.Println(*b) // *b gets the value stored at the memory address stored in the b i.e., 42 (a=42)
	//This is called dereferencing

	// b is a pointer to the memory address where an int is stored
	// b is a type of int pointer
	x := 5
	zero(&x) //passing the address of the x
	fmt.Println(x)
	//this will be returned 0 because the zero function changes the value available at the memory address of x

	ab := 5
	zeroa(ab)
	fmt.Println(ab)
	//the value of ab will still be 0 because the memory addresses are different and scope as well

}

func zero(y *int) {
	*y = 0 //changing value available at the address of y
}
func zeroa(aa int) {
	aa = 0
}
