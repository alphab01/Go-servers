package main

import (
  "fmt"
  "net/http"
  "net/url"
  "io"
)

func main() {
  bu := "https://example.com/api/resource"
  p := url.Values{}
  p.Add("p1", "v1")
  p.Add("p2", "v2")
  url := bu + "?" + p.Encode()
  resp, err := http.Get(url)
  if (err != nil) {
    fmt.Println(err)
    return
  }
  defer resp.Body.Close()
  d, err := io.ReadAll(resp.Body)
  if (err != nil) {
    fmt.Println(err)
    return
  }
  fmt.Printf("%s", d)
}
