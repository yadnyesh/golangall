package main

import "fmt"

func main() {
	//s := greet ("Jane", "Done")
	fmt.Println(greet ("Jane", "Done"))
}

func greet(fname string, lname string) string {
	return fmt.Sprint(fname, lname)
}