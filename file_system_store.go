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

func (f *FileSystemPlayerStore) GetPlayerScore(name string) (wins int) {
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}
	return
}
