package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)


func main() {

	res, err := http.Get("https://www.google.co.in")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)

	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", page)

}

