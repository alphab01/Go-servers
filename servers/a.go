package main

import (
  "net/http"
)

func h(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("hello world"))
}

func main() {
  http.HandleFunc("/", h)
  _ = http.ListenAndServe(":8080", nil)
}
