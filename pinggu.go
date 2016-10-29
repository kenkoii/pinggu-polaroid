package main

import (
        "net/http"
)

func main() {
    http.HandleFunc("/", static)
}

func static(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "public/"+r.URL.Path)
}