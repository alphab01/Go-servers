package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io"
  "net/http"
)

type usr struct {
  Name string `json:"name"`
  ID uint32 `json:"id"`
}

func main() {
  var s = usr{Name: "Kirill", ID: 1,}
  br, err := json.Marshal(s)
  if (err != nil) {
    fmt.Println(err)
    return
  }
  r, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(br))
  if (err != nil) {
    fmt.Println(err)
    return
  }
  br2, err := io.ReadAll(r.Body)
  if (err != nil) {
    fmt.Println(err)
    return
  }
  fmt.Println(string(br2))
}
