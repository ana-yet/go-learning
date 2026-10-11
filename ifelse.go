package main

// if else with scope
func main() {
	age := 60
	println(age)

	// if else with scope || we can declare a variable in the if statement and use it in the else block
	if age := 50; age >= 60 {
		println("old citizen", age)
	} else if age >= 60 {
		println("User is adult", age)
	} else {
		println("user is under age", age)
	}
}