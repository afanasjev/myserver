package main

import (
    "log"
    "net/http"
)

type inMemory struct {}

func (i *inMemory) GetPlayerScore(name string) int {
    return 123
}

func main() {
    
    server := &PlayerServer{store: &inMemory{}}
    if err := http.ListenAndServe("0.0.0.0:8888", server); err != nil {
        log.Fatalf("could not listen on port 8888 because of %v", err)
    }
}
