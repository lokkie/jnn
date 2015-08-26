package main
import (
    "github.com/gorilla/mux"
)

type (
    JnnRouter struct {
        *mux.Router
    }    
)

func (mr *JnnRouter) Installroutes() {
    mr.HandleFunc("/api/me/put", putMe)
}