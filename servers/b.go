package main

import (
  "fmt"
  "net/http"
)

func h(w http.ResponseWriter, r *http.Request) {
  fmt.Println(r.URL.String())
  w.Write([]byte("hello, world"))
}

func main() {
  http.HandleFunc("/", h)
  err := http.ListenAndServe(":8080", nil)
  if (err != nil) {
    fmt.Println(err)
    return
  }
}
