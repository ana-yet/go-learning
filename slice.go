package main

import "fmt"

func main() {
	var order = [6]int{1, 2, 3, 4, 5, 6}

	slice := order[1:4] // create a slice from the array

	fmt.Println(slice)

	// capacity is the number of elements in the underlying array, starting from the first element of the slice and length is the number of elements in the slice itself

	// capacity of the slice
	fmt.Println(cap(slice))

	// length of the slice
	fmt.Println(len(slice))

	// slice len and capacity can be changed by appending elements to the slice
	slice = append(slice, 7)
	fmt.Println(slice) // output: [2 3 4 7]
	fmt.Println(cap(slice)) // output: 5
	fmt.Println(len(slice)) // output: 4
}