package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println("hello http server")
	})

}
