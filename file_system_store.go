package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() (league []Player) {
	f.database.Seek(0, 0)
	league, _ = NewLeague(f.database)
	return
}
