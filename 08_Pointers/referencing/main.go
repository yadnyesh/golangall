package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a);
	fmt.Println(&a);

	var b *int = &a //referencing

	fmt.Println(b)
	fmt.Println(&b)
	fmt.Println(*b) //de-referencing

	*b = 42
	fmt.Println(a);

}
