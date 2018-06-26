package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err = http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
		os.exit(1)
	}
	fmt.Println(resp)
}
