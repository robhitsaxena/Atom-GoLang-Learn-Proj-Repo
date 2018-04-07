package main

import "fmt"

func main() {
	i := 0
	for {
		i++

		if i%2 == 0 {
			fmt.Printf("%s-%d\n", "Printing Even", i)
			continue // Will jump back to the start of for loop, incrementing i++
		}
		fmt.Printf("%s-%d\n", "Printing Odd", i)
		if i >= 10 {
			break // will put us out of the loop
		}

	}
}
