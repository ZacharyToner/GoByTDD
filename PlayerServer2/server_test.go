package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

func TestGetPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  30,
		},
		nil,
		nil,
	}

	server := NewPlayerServer(&store)

	t.Run("returns Pepper's Score", func(t *testing.T) {
		req := newGetScoreRequest("Pepper")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertResponseCode(t, resp.Code, http.StatusOK)
		assertResponseBody(t, resp.Body.String(), "20")
	})

	t.Run("returns Floyd's Score", func(t *testing.T) {
		req := newGetScoreRequest("Floyd")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertResponseCode(t, resp.Code, http.StatusOK)
		assertResponseBody(t, resp.Body.String(), "30")

	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		req := newGetScoreRequest("Tom")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertResponseCode(t, resp.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}

	server := NewPlayerServer(&store)

	t.Run("it records a win on POST", func(t *testing.T) {
		player := "Apollo"
		req := newPostWinRequest(player)
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertResponseCode(t, resp.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("Function call mismatch: got %d calls to RecordWin, want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("Func Call Argument Missmatch: Called RecordWin for '%s' but should have been '%s'", store.winCalls[0], player)
		}
	})
}

func TestRecordingWinsAndGettingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		resp := httptest.NewRecorder()
		server.ServeHTTP(resp, newGetScoreRequest(player))
		assertResponseCode(t, resp.Code, http.StatusOK)
		assertResponseBody(t, resp.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		resp := httptest.NewRecorder()
		server.ServeHTTP(resp, newLeagueRequest())
		assertResponseCode(t, resp.Code, http.StatusOK)
		got := getLeagueFromResponse(t, resp.Body)
		want := []Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})

}

func TestLeague(t *testing.T) {
	t.Run("It returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		req := newLeagueRequest()
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		got := getLeagueFromResponse(t, resp.Body)

		assertResponseCode(t, resp.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, resp, jsonContentType)
	})
}

func newGetScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func newPostWinRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", body, err)
	}
	return
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Response Body Mismatch: got '%s' but wanted '%s'", got, want)
	}
}

func assertResponseCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Response Code Mismatch: got %d, want %d", got, want)
	}
}

func assertLeague(t *testing.T, got, want []Player) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("League Equality Mismatch: got %v, wanted %v", got, want)
	}
}

func assertContentType(t *testing.T, resp *httptest.ResponseRecorder, want string) {
	t.Helper()
	if resp.Header().Get("content-type") != want {
		t.Errorf("response did not have content-type of '%s', got %v", want, resp.HeaderMap)
	}
}
