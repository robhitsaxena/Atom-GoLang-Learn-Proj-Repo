package main

import "fmt"

func main() {
	n := average(2, 3, 4, 5, 6)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64 //this will set total to a value 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))

}
