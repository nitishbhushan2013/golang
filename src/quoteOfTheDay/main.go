package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://quotes.rest/qod.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(string(content))
}
