package main

import (
	"fmt"
)

var x = 0

func main() {

	//x := 0
	//increment := func() int {
	//	x++
	//	return x
	//}

	fmt.Println(increment())
	fmt.Println(increment())
}

func increment() int {
	x++
	return x
}
