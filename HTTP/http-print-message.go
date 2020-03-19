package main

import (
	"fmt"
	"net/http"
)

type listener int //implementing the handler interface

func (m listener) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "my code from vishu") //takes 2nd argument as message and writes to 'w' writer and prints the message
}

func main() {
	var ls listener
	http.ListenAndServe(":8080", ls) //takes the port number (listens to it) and the handler variable
	//It creates the small server and executes the given operations on it.
}

//Run this program and visit "http://localhost:8080" on browser.
