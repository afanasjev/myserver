package main

import (
    "net/http"
    "fmt"
    "strings"
)

type PlayerServer struct {
    store PlayerStore
    http.Handler
}

type PlayerStore interface {
    GetPlayerScore(string) int
    RecordWin(string)
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
    server := &PlayerServer{ store: store }

    router := http.NewServeMux() 

    router.Handle("/league", http.HandlerFunc(server.leagueHandler))
    router.Handle("/players/", http.HandlerFunc(server.playerHandler))
    server.Handler = router

    return server
}

func (p *PlayerServer) leagueHandler (w http.ResponseWriter, r *http.Request){
    w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playerHandler (w http.ResponseWriter, r *http.Request){
    player := strings.TrimPrefix(r.URL.Path, "/players/")
    switch r.Method {
    case http.MethodPost:
        p.processWin(w, player)
    case http.MethodGet:
        p.showScore(w, player)
    }
}

func (p *PlayerServer) processWin(response http.ResponseWriter, player string) {
    p.store.RecordWin(player)
    response.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(response http.ResponseWriter, player string) {
    score := p.store.GetPlayerScore(player)
    if score == 0 {
        response.WriteHeader(http.StatusNotFound)
    }
    fmt.Fprint(response, score)
}
