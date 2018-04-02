package Variables

import (
	"fmt"
	"strconv"
)

func PrintVariables() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	var e string
	var f int
	var g float64
	var h bool

	fmt.Println(fmt.Sprintf("%s:%d", "Value of a is: ", a))
	fmt.Println("Value of b is: " + b)
	fmt.Println(fmt.Sprintf("%s:%g", "Value of c is: ", c))
	fmt.Println("Value of d is: " + strconv.FormatBool(d))

	fmt.Printf("Type of a is: "+"%T \n", a)
	fmt.Printf("Type of b is: "+"%T \n", b)
	fmt.Printf("Type of c is: "+"%T \n", c)
	fmt.Printf("Type of d is: "+"%T \n", d)

	fmt.Printf("Type of e is: "+"%v \n", e)
	fmt.Printf("Type of f is: "+"%v \n", f)
	fmt.Printf("Type of g is: "+"%v \n", g)
	fmt.Printf("Type of h is: "+"%v \n", h)

}
