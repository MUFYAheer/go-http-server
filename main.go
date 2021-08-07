package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("could not read file, %v", err)
	}
	defer db.Close()

	store := NewFileSystemPlayerStore(db)

	server := NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":9000", server))
}
