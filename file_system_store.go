package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func (f *FileSystemPlayerStore) GetLeague() (league []Player) {
	json.NewDecoder(f.database).Decode(&league)
	return
}
