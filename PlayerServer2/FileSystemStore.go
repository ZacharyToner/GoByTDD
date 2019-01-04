package main

import (
	"encoding/json"
	"io"
)

// FileSystemPlayerStore is a PlayerStore implementation using the filesystem
type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

// GetLeague reads data from the database, ensuring to go back to the beginning
func (f *FileSystemPlayerStore) GetLeague() League {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

// GetPlayerScore returns the score of the requested player
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

// RecordWin increments number of wins for requested player
func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}
