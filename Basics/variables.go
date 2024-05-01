package main

import "fmt"

func main() {

	//way 1 declaration and definition
	var a, b, c, d int = 1, 2, 3, 4
	fmt.Println("a: ", a, "\nb: ", b, "\nc: ", c, "\nd: ", d)

	//way 2 declaration and definition
	var e, f = 1, "hindi"
	g, h := 2, "loyal"

	fmt.Println("\ne: ", e, "\nf: ", f, "\ng: ", g, "\nh: ", h)

	//way 3 declaration and definition
	var (
		p int    //default value is 0 printed
		q int    = 10
		r string = "hello"
	)

	fmt.Println("\np: ", p, "\nq: ", q, "\nr: ", r)
	const pi = 3.14
	fmt.Println(pi)
}
