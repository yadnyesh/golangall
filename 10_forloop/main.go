package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, " - ", j)
		}
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	///don't do
	i = 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break;
		}
		i++
	}

	for j:=25000 ; j < 25100; j++ {
		fmt.Println(i, " - ", string(j), " - ", []byte (string(j)))
	}



}
