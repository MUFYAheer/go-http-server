package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())

	log.Fatal(http.ListenAndServe(":9000", server))
}
