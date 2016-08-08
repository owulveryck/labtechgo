package main

import (
	"log"
	"net/http" // HL
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Println("Listening on 8000")
	http.ListenAndServe(":8000", nil)
}
