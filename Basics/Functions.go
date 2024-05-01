package main

import "fmt"

func solve3(a int, b string) (res1 int, res2 string) {
	res1 = a + a
	res2 = b + " World"

	//fmt.Println(a, b)

	return
}

func solve2(arr []int) int {
	res := 0
	for _, v := range arr {
		res += v
	}
	return res
}

func solve() {
	//fmt.Println("Simple function")

	n := 0
	fmt.Scan(&n)
	//var arr = []int {}
	arr := make([]int, n)
	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}
	res := solve2(arr)
	fmt.Println(res)

	a, b := solve3(5, "Hello")
	fmt.Println(a, b)

	//only a
	x, _ := solve3(6, "Magical")
	fmt.Println(x)
}

func main() {

	//Functions
	//Simply create a function with other name than main outside main function and call in main function
	t := 0
	fmt.Scan(&t)
	for i := 0; i <= t; i++ {
		solve()
	}

}
