package main

import (
	"fmt"
	"net/http"
)

func startServer() {
	http.HandleFunc("/", helloUser)
	http.ListenAndServe(":8080", nil)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user! Welcome to the to do list app!"
	fmt.Fprintln(writer, greeting)
}
