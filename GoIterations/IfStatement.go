package GoIterations

import "fmt"

func GoIfStatement() {
	b := true
	if food := "chocolate"; b {
		fmt.Println(food)
	}
	if name := "Robhit"; name == "Robhits" {
		fmt.Println("Correct name!")
	} else {
		fmt.Println("Wrong Name!")
	}
}
