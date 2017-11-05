package main

import "fmt"

func main() {
	name := "Yadnyesh"
	fmt.Print("Please enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello, ", name)
}
