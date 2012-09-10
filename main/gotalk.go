package main

import (
	"gotalk/gotalk"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", gotalk.Router)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
