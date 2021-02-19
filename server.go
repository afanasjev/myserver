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
    RecordWin(string)
}

func (p *PlayerServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    router := http.NewServeMux()

    router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        w.WriteHeader(http.StatusOK)
    }))

    router.Handle("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        player := strings.TrimPrefix(r.URL.Path, "/players/")
        switch r.Method {
        case http.MethodPost:
            p.processWin(w, player)
        case http.MethodGet:
            p.showScore(w, player)
        }
    }))

    router.ServeHTTP(response, request)
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
