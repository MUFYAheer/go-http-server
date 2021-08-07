package main

import (
	"log"
	"net/http"
	"os"
	"poker"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("could not read file, %v", err)
	}
	defer db.Close()

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	server := poker.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":9000", server))
}
