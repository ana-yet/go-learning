package main

import "fmt"

func main() {

	// anonymous function
	coffeemaker := func() {
		fmt.Printf("coffee ready guys \n");
	}

	coffeemaker()

	// IIFE (Immediately Invoked Function Expression)
	func(name string) {
		fmt.Printf("%s coffee ready guys \n",name);
	}("Black")

	// block scope 
	{ 
		// under this block scope, the variable is only accessible within this block
		name := "anayet"
		fmt.Printf("my name is %s \n",name);
	}


	// variable scope
	name := "anayet"
	{
		name := "name in block scope"
		fmt.Printf("my name is %s \n",name); // output: my name is name in block scope
	}
	fmt.Printf("my name is %s \n",name); // output: my name is anayet
	
	
}


