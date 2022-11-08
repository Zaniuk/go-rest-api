package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	log.Println("Server running on port 80")
	http.ListenAndServe(":80", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo!")
}
