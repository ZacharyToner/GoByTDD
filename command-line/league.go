package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

// League is a more contextual representation of []Player
type League []Player

// Find pulls the requested player from the []Player
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

// NewLeague generates a slice of Players by parsing the Reader as JSON
func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return league, err
}
