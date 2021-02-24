package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "fmt"
    "encoding/json"
)

type StubPlayerStore struct {
    scores map[string]int
    winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
    score := s.scores[name]
    return score
}

func (s *StubPlayerStore) RecordWin(name string) {
    s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
    store := StubPlayerStore{scores: map[string]int{
        "Pepper"    : 20,
        "Floyd"     : 10,
    }}
    server := NewPlayerServer(&store)

    t.Run("return Pepper's score", func(t *testing.T){
        request := newGetScoreRequest("Pepper")
        response := httptest.NewRecorder()
        server.ServeHTTP(response, request)
        assertStatus(t, response.Code, http.StatusOK)
        assertResponseBody(t, response.Body.String(), "20")
    })

    t.Run("return Floyd's score", func(t *testing.T){
        request := newGetScoreRequest("Floyd")
        response := httptest.NewRecorder()
        server.ServeHTTP(response, request)
        assertStatus(t, response.Code, http.StatusOK)
        assertResponseBody(t, response.Body.String(), "10")
    })
    t.Run("return Floyd's score", func(t *testing.T){
        request := newGetScoreRequest("PlayerNotExists")
        response := httptest.NewRecorder()
        server.ServeHTTP(response, request)

        assertStatus(t, response.Code, http.StatusNotFound)
    })
    t.Run("wins on POST", func(t *testing.T){
        player := "bob"
        request := newPostWinRequest(player)
        response := httptest.NewRecorder()
        server.ServeHTTP(response, request)

        assertStatus(t, response.Code, http.StatusAccepted)
        got := len(store.winCalls)
        if  got != 1 {
            t.Errorf("got %d want %d\n", got, 1)
        }

        if store.winCalls[0] != player {
            t.Errorf("got %q want %q\n", store.winCalls[0], player)
        }
    })
}

func TestLeague(t *testing.T){
    store := StubPlayerStore{}
    server := NewPlayerServer(&store)

    t.Run("/league returns 200 OK", func(t *testing.T){
        request, _ := http.NewRequest(http.MethodGet, "/league", nil)
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)
        var got []Player

        err := json.NewDecoder(response.Body).Decode(&got)

        if err != nil {
            t.Errorf("Unable decode json %q with error: %v\n", response.Body, err)
        }

        assertStatus(t, response.Code, http.StatusOK)
    })

}

func newGetScoreRequest(player string) *http.Request {
    path := fmt.Sprintf("/players/%s", player)
    request, _ := http.NewRequest(http.MethodGet, path, nil)
    return request
}

func assertResponseBody(t *testing.T, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q\n", got, want)
    }
}

func assertStatus(t *testing.T, got, want int) {
    t.Helper()
    if got != want {
        t.Errorf("got %d want %d\n", got, want)
    }
}

func newPostWinRequest(name string) *http.Request {
    request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
    return request
}
