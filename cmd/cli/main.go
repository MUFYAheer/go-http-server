package main

import (
	"fmt"
	"log"
	"os"
	"poker"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Fatalf("problem opening %s, %v", dbFileName, err)
	}
	playerStore, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem initializing file system player store, %v", err)
	}

	game := poker.NewCLI(playerStore, os.Stdin)
	game.PlayPoker()
}
