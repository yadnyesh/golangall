package main

import "fmt"

func hello() {
	fmt.Println("Hellow")
}

func world() {
	fmt.Println("World")
}


func main() {
	defer world()
	hello()
}