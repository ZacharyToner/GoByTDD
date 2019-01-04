package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndGettingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "")
	defer cleanDatabase()
	store := NewFileSystemPlayerStore(database)
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
