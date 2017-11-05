package main

import (
	"fmt"
	"strconv"
)


func main() {

	const q  = "ABC"
	const (
		Pi = 3.14
		Language = "Go"
	)
	//res, _ := http.Get("https://www.google.co.in")
	//
	//page, _ := ioutil.ReadAll(res.Body)
	//
	//res.Body.Close()
	//
	//fmt.Printf("%s", page)
	fmt.Printf(strconv.FormatFloat(Pi, 'g', 1, 64))
	fmt.Printf(Language)
	fmt.Printf(q)

}

