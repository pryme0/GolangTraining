package main

import (
	"net/http"
)

func main() {
	// parse csv

	// parse template

	// function
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// execute template

	})

	// create server
	http.ListenAndServe(":9000", nil)
}
