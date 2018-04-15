package GoIterations

import "fmt"

//We are creating own (user defined) type here

type Contact struct {
	Greeting string
	Name     string
}

func GoOnTypes(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an Int type")
	case string:
		fmt.Println("It's a String")
	case Contact:
		fmt.Println("It's an user defined type - contact")
	default:
		fmt.Println("Unknown Type")
	}

}
