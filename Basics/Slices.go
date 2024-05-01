package main

import "fmt"

func main() {

	//similar way as array, size not defined 
	//Unlike Array, size of slice can be modified, see below
 
	var slice1 = []string{"Hello", "Dear", "friend"}
	fmt.Println(slice1)

	//append slices
	var slice2 = []string{"Hi", "dude"}

	slice3 := append(slice1, slice2...)

	fmt.Println(slice3)

	slice3 = append(slice3, "Aman", "Kishor")

	fmt.Println(slice3)

	arr := [5]int{1, 2, 3, 4, 5}

	// slice the array
	slice4 := arr[1:4]
	fmt.Println(slice4)

}
