package main

import "fmt"

func main() {
	//print decimal
	fmt.Println(42)

	//print binary
	fmt.Printf("%d - %b\n", 42, 42)

	//print Hexadecimal
	fmt.Printf("%d - %#x\n", 42, 42)

	for i := 0; i < 300; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}

