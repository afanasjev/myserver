package main

import (
    "log"
    "net/http"
)

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
    return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
    store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
    i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
    return i.store[name]
}

func main() {
    
    server := NewPlayerServer(NewInMemoryPlayerStore())
    if err := http.ListenAndServe("0.0.0.0:8888", server); err != nil {
        log.Fatalf("could not listen on port 8888 because of %v", err)
    }
}
