package main

func main() {
	fruit := []string{"apple", "banana", "orange", "grapes"}

	fruit[0] = "mango"                      // change the value of the first element of the array
	fruit = append(fruit, "kiwi")           // add a new element to the array
	fruit = append(fruit[:2], fruit[3:]...) // remove the third element of the array

	println(fruit)      // output: [mango banana grapes kiwi]
	println(len(fruit)) // output: 4
	println(fruit[0])   // output: mango

	// loop through the array using range
	for i, v := range fruit {
		println(i, v)
	}

	// loop through the array using traditional for loop
	for i := 0; i < len(fruit); i++ {
		println(i, fruit[i])
	}
}