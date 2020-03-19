package main

import (
	"fmt"
	"log"
	"net/http"
)

var HitCounter int

func main() {
	http.HandleFunc("/users", usersHandleFunc) //it registers the handler function for working of HTTP
	log.Fatal(http.ListenAndServe(":8080", nil)) //starts, displays and listens on the server
}

func usersHandleFunc(w http.ResponseWriter, r *http.Request) {
	HitCounter++
	fmt.Printf("Received a request '%v' times.\n",HitCounter) //will be displayed on the console/terminal
	fmt.Fprintf(w, "Hello there. This was received using method '%v' and you called this page '%v' times.", r.Method,HitCounter)
}

//Run this program and visit "http://localhost:8080/users" on browser.
//Whenever you visit that URL and program is running, the message "Received a request." will be displayed on console/terminal.
