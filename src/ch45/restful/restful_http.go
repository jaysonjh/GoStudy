package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Welcome!")
}

func Hello(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintln(w, "Hello, %s!/n", params.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	http.ListenAndServe(":8080", router)
}
