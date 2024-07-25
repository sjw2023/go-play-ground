package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	MyRouter := mux.NewRouter()

	MyRouter.HandleFunc("/{title}", func(resp http.ResponseWriter, req *http.Request) {
		VarsInRequest := mux.Vars(req)
		fmt.Println(VarsInRequest["title"])
	})

	MyRouter.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println("hello http server")
	})

	http.ListenAndServe(":8080", MyRouter)
	// http.ListenAndServe(":8080", nil)
}
