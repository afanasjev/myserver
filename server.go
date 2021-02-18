package main

import (
    "net/http"
    "fmt"
    "strings"
)

type PlayerServer struct {
    store PlayerStore
}

type PlayerStore interface {
    GetPlayerScore(string) int
}

func (p *PlayerServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    player := strings.TrimPrefix(request.URL.Path, "/players/")
    score := p.store.GetPlayerScore(player)
    if score == 0 {
        response.WriteHeader(http.StatusNotFound)
    }
    fmt.Fprint(response, score)
}
