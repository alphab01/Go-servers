package main

import (
  "fmt"
  "net/http"
  "io"
)

func main() {
  r, _ := http.Get("http://localhost:8080")
  defer r.Body.Close()
  d, _ := io.ReadAll(r.Body)
  fmt.Println("%s", d)
}
