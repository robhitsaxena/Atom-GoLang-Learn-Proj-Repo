package main

import "fmt"

func main() {
	fmt.Println(greet("Robhit", "Saxena"))
	fmt.Println(greetAnother("Robhit", "Saxena"))
	fmt.Println(MultiVal("Robhit", "Saxena"))
}

func greet(fname string, lname string) string /* Here we are returning a string */ {
	return fmt.Sprint(fname, " ", lname)
	//Sprint is string print.

}

func greetAnother(fname string, lname string) (s string) /* Here s is going to be returned */ {
	s = fmt.Sprint(fname, " ", lname)
	return

}

// Return multiple values

func MultiVal(fname, lname string) (string, string) {
	return fmt.Sprint(fname, " ", lname), fmt.Sprint(lname, " ", fname)

}
