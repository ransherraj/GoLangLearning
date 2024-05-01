package main

import "fmt"

func main() {

	//Array
	//For loop
	//range keyword : to get index as well as value
	//range can be used on array, slice, map as syntax below

	x := 5

	for i := 0; i < x; i++ {
		fmt.Printf("Index %v at Place %v", i, i+1)
		fmt.Println()
	}

	taste := []string{"Sweet", "Sour"}
	fruit := []string{"Mango", "Orange", "Apple"}

	for i := 0; i < len(taste); i++ {
		for j := 0; j < len(fruit); j++ {
			fmt.Println(taste[i], " ", fruit[j])
		}
	}

	//range
	for i, v := range fruit {
		fmt.Printf("At index %v value in fruit array is %v ", i, v)
		fmt.Println()
	}

	//to ommit index use _

	for _, v := range fruit {
		fmt.Printf("%v ", v)
		fmt.Println()
	}

}
