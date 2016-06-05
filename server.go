package main

import "net/http"

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        w.Write([]byte("Welcome to the home page!"))
    })
    http.ListenAndServe(":8080", mux)
}

