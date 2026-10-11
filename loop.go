// go has only one loop statement, the for loop. The for loop can be used in several ways.

package main

func main(){
	// for loop with only condition
	// for i :=0 ; i < 10; i++ { // for initialization; condition; post statement
	// 	fmt.Println(i)
	// }
	
	// while style for loop
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++  // increment the value of i by 1, if we don't increment the value of i, it will be an infinite loop
	// }

	// infinite loop 
	// for {
	// 	fmt.Println("infinite loop")
	// }

	// break statement
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break // break statement will exit the loop
	// 	}
	// 	fmt.Println(i)
	// }

	// continue statement
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		continue // continue statement will skip the current iteration and move to the next iteration
	// 	}
	// 	fmt.Println(i)
	// }
}