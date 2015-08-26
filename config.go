package main

import (
    "encoding/json"
    "io/ioutil"
)

type Configuration struct {
    Database struct {
        Host        string `json:"host"`
        Port        int `json:"port"`
        User        string `json:"user-name"`
        Password    string `json:"access-password"`
        Name        string `json:"database-name"`
    } `json:"database"`
}

func (conf *Configuration) Load(path string) error {
    jsonSrc, e := ioutil.ReadFile(path)
    if e != nil {
        return e
    }
    
    json.Unmarshal(jsonSrc, &conf)
    return nil
}
