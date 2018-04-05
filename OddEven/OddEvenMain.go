package main

import "fmt"

func main() {
	OddEven()
}
func OddEven() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Printf("%s%d\n", "Number is Odd - ", i)
		} else {
			fmt.Printf("%s%d\n", "Number is Even - ", i)
		}
	}
}
