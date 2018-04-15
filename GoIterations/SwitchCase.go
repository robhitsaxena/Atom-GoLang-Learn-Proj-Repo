package GoIterations

import "fmt"

func SwitchCaseMain() {
	switch "Robhit" {
	case "Robhit", "Saxena": //Multiple case conditions.
		fmt.Println("You are the right person!")
		fallthrough //If this case is true then run the next case condition.
	case "Rikku":
		fmt.Println("You are the right person!")
	default:
		fmt.Println("No Match")

	}
}
