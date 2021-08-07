package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func NewLeague(reader io.Reader) (league []Player, err error) {
	err = json.NewDecoder(reader).Decode(&league)
	if err != nil {
		err = fmt.Errorf("error parsing league, %v", err)
	}
	return
}
