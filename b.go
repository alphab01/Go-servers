package main

import (
  "fmt"
  "io"
  "net/http"
)

func main() {
  r, err := http.Get("https://golang.org")
  if (err != nil) {
    fmt.Println(err)
    return
  }
  defer r.Body.Close()
  d, err := io.ReadAll(r.Body)
  if (err != nil) {
    fmt.Println(err)
    return
  }
  fmt.Printf("%s", d)
}
