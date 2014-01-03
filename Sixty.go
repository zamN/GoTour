package main

import (
    "fmt"
    "net/http"
)

type String string

type Msg struct {
    Greeting string
    Punct string
    Who string
}

func  (m *Msg) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, m.Greeting, m.Punct, m.Who)
}

func  (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, s)
}

func main() {
    http.Handle("/msg", &Msg{"Hello", ":", "Gophers!"})
    http.Handle("/str", String("Hello friend."))
    http.ListenAndServe("localhost:4000", nil)
}
