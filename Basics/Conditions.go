package main

import "fmt"

func main() {

	//if/ else if/ else
	//Error occurs if 'else if' or 'else' is applied
	//to the next line of curly braces of if

	// This is incorrect

	// if a < b {
	// 	fmt.Println("a < b")
	// }
	// else if a > b {
	// 	fmt.Println("a > b")
	// }
	// else{
	// 	fmt.Println("a = b")
	// }

	a := 10
	b := 10
	if a < b {
		fmt.Println("a < b")
	} else if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a = b")
	}

}
