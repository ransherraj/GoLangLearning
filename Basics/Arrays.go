package main

import "fmt"

func main() {

	//way 1
	var arr = [5]int{1, 2, 3}

	//way 2
	var strarr = [5]string{"Hello", "Dear", "friend"}
	const pi = 3.14
	fmt.Println(arr, "size of arr:", len(arr))
	fmt.Println(strarr,"size of strarr:", len(strarr))

	//way 3
	var newArr = [...]int{1, 2, 3, 5, 7}
	newArr2 := [...]int{10, 202, 30, 40, 100, 60, 70}

	//Modification
	newArr[2] = 20000
	strarr[2] = "Raj"
	fmt.Println(newArr, "size of newArr:", len(newArr))
	fmt.Println(newArr2,"size of newArr2:", len(newArr2))
	fmt.Println("new modified strarr: ", strarr,"size of strarr: ", len(strarr))

	//Partial assign
	partialArr := [6]string{3:"raj", 5:"Anmol"}
	fmt.Println(partialArr,"size of partialArr:", len(partialArr))
}
