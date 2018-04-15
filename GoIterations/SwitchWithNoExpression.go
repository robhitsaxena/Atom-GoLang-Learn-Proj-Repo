package GoIterations

import "fmt"

func SwitchNoExp() {
	myFriendName := "Robhit"

	switch {
	case len(myFriendName) == 6:
		fmt.Println("Matches the first one")
	case myFriendName == "Robhit":
		fmt.Println("This matches bu won't be returned because switch with no expression got the match in first the case")
	default:
		fmt.Println("Nothing matches, printing default!")

	}
}
