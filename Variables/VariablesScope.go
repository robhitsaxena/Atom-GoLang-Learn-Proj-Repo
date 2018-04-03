package main

import "fmt"

var a int = 0
var b int = 0

func increment() int {
	a++
	return a
}

//Returning a function from within a Function.
//Here the return type is function with int.
func printFunc() func() int {
	as := 19
	return func() int {
		return as
	}
}
func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(fmt.Sprintf("%s:%d", "Accessing Variable Under Inner Scope", x))
		fmt.Println(fmt.Sprintf("%s:%d", "Global Variable Scope Example", a))
		fmt.Println(fmt.Sprintf("%s:%d", "Global Variable Scope Example", increment()))
		fmt.Println(increment())
		fmt.Println(increment())
		fmt.Printf("Type of a is: "+"%T \n", a)

		//Function Expression & Anonymous Function
		//It is used to put function inside a fucntion otherwise you will hit an error.
		incrementVar := func() int {
			b++
			return b
		}
		fmt.Println(fmt.Sprintf("%s:%d", "Function inside a function", incrementVar()))

	}

	//returning function within a fucntion
	fmt.Println(printFunc())
	fmt.Printf("Type of a is: "+"%T \n", printFunc())
	//The above will return the memory location where the value is stored.
	//Something like a pointer.
	//The below code will give the value because we have to assign the function return value to something.
	varPrintFunc := printFunc()
	fmt.Println(varPrintFunc())
	fmt.Printf("Type of a is: "+"%T \n", varPrintFunc())

}
