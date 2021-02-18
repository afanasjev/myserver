package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "fmt"
)

type StubPlayerStore struct {
    scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
    score := s.scores[name]
    return score
}

func TestGETPlayers(t *testing.T) {
    store := StubPlayerStore{map[string]int{
        "Pepper"    : 20,
        "Floyd"     : 10,
    }}
    server := &PlayerServer{&store}

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
