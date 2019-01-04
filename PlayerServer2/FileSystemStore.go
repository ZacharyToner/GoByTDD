package main

import (
	"encoding/json"
	"io"
)

// FileSystemPlayerStore is a PlayerStore implementation using the filesystem
type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
	league   League
}

// NewFileSystemPlayerStore creates a new Store, allowing for only one nevessary read
func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		database: database,
		league:   league,
	}
}

// GetLeague reads data from the database, ensuring to go back to the beginning
func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

// GetPlayerScore returns the score of the requested player
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

// RecordWin increments number of wins for requested player
func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(f.league)
}
