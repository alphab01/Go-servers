package main

import (
  "fmt"
  "encoding/json"
  "net/http"
  "net/url"
)

func main() {
  form := url.Values{
    "name": {"hello"},
    "sname": {"golang post"},
  }
  r, err := http.PostForm("https://httpbin.org/post", form)
  if (err != nil) {
    fmt.Println(err)
    return
  }
  var res map[string]interface{}
  json.NewDecoder(r.Body).Decode(&res)
  fmt.Println(res["form"])
}
