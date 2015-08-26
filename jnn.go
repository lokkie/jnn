package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "time"
)


func main() {
    c := &Configuration{}
    e := c.Load("./conf/main.json") 
    if e == nil {
        }
    s := Startservice()
    s.Installroute();
    
    fmt.Printf("Go-go-go\n")
}


type Todo struct {
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
}

type Todos []Todo

func putMe(w http.ResponseWriter, r *http.Request) {
    todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }    
}