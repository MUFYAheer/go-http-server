package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

func NewLeague(reader io.Reader) (league League, err error) {
	err = json.NewDecoder(reader).Decode(&league)
	if err != nil {
		err = fmt.Errorf("error parsing league, %v", err)
	}
	return
}
