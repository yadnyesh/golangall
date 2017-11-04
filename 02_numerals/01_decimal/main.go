package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)


func main() {

	res, _ := http.Get("https://www.google.co.in")

	page, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()

	fmt.Printf("%s", page)

}

