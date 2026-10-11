package main

import "fmt"

func sum (a int, b int) int {
	return a + b
}

func main() {
	// name := "anayet" short variable declaration

	// group variable declaration
	var (
		name string = "anayet";
		age int = 20;
	)

	// multiple variable declaration
	var a, b, c int = 1, 2, 3;

	const pi float64 = 3.14 // constant declaration
	// pi = 3.14159 // this will cause an error because constants cannot be reassigned


	// invalid variable names
	// userName := "anayet" // variable name with camelCase
	// _user := "anayet" // variable name starting with underscore

	// output the values of the variables

	// group variable declaration
	fmt.Println(name, age)

	// multiple variable declaration
	fmt.Println(a,b,c)

	// placeholder variable declaration
	var cName string = "company name"
	fmt.Printf("what is your %s? \n", cName)

	// call function
	var result = sum(20,30)
	fmt.Println(result)
}
