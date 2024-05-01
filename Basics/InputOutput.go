
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main(){
	// var x, y int
	// var z string
	// var a bool
	// var b float32
	// var c int32
	var e rune

	// fmt.Println("Enter x and y:")
	// fmt.Scan(&x, &y);

	// fmt.Println("Enter z:")
	// fmt.Scan(&z);

	// fmt.Println("Enter a, b, c:")
	// fmt.Scan(&a, &b, &c);

	fmt.Println("Enter e:")
	// _, err := fmt.Scanf("%c", &e)
	reader := bufio.NewReader(os.Stdin);
	e,_,err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("sum of x and y is : ", x+y);
	// fmt.Println("z is : ", z);
	// fmt.Println("a : ", a, "b : ", b, "c : ", c, "e : ", e);
	fmt.Println("e : ", e);
}