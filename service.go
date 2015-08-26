package main
import (
    "github.com/gorilla/mux"
    "net/http"
    "log"
)

type (
    Service struct {
        Router *JnnRouter
    }
    
)

func Startservice() (*Service) {
    return &Service{}
}

func (s *Service) Installroute() {
    s.Router = &JnnRouter{mux.NewRouter().StrictSlash(true)}
    s.Router.Installroutes()
    log.Fatal(http.ListenAndServe(":8080", s.Router))
}