package main

//InMemoryPlayerStore is an in-memory implementation of the PlayerStore interface
type InMemoryPlayerStore struct {
	store map[string]int
}

//NewInMemoryPlayerStore is a constructor for InMemeoryPlayerStore
//InMemoryPlayerStore utilizes a map
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

//GetPlayerScore returns stored value from map for requested name
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

//RecordWin increments value in map for requested name
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

//GetLeague returns a representation of all players in the store
func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player

	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}

	return league
}
