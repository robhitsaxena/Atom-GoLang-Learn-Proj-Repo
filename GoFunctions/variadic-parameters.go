package main

import "fmt"

func main() {
	data := []float64{1, 2, 3}
	n := average(data...) //This is variadic parameters where ... is in the last. Means passing the list data one by one.
	fmt.Println(n)
	d := alternateAverage(data)
	fmt.Println(d)

}

func average(sf ...float64) float64 {
	//println(sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func alternateAverage(sf []float64) float64 {
	//println(sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
