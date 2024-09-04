package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "<html><h1>Hello %s!</h1></html>",req.URL.Path[1:])
}

func main() {

    http.HandleFunc("/", hello)

    http.ListenAndServe(":8090", nil)
}
