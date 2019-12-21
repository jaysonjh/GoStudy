package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello world")
	})

	http.HandleFunc("/time", func(writer http.ResponseWriter, request *http.Request) {
		time := time.Now()
		timeStr := fmt.Sprintf("{\"time\" : \" %s\" }", time)
		writer.Write([]byte(timeStr))
	})

	http.ListenAndServe(":8080", nil)
}
