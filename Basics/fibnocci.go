package main

import "fmt"



func fibnocci(n int) ([]int) {
	arr := make([]int, n+1)
	// if n == 0 || n == 1 {
	// 	arr[n] = n
	// }
	if n > 0{
		arr[0] = 0
	}
	if n > 1{
		arr[1] = 1
	}
	
	for i:=2; i<=n; i++{
		arr[i] = arr[i-1] + arr[i-2];
		//fmt.Println(arr[i])
	}
	//fmt.Println(arr)
	return arr
}

func solve6() {
	n := 0
	fmt.Scan(&n)
	res := fibnocci(n);
	fmt.Println(res)
}

func main() {

	//Functions
	//Simply create a function with other name than main outside main function and call in main function
	t := 0
	fmt.Scan(&t)
	for i := 0; i <= t; i++ {
		solve6()
	}

}
