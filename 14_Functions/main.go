package main

import "fmt"

func main() {
	fmt.Println(greet ("Jane", "Done"))
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(fname, lname)
}