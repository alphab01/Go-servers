package main

import (
  "fmt"
  "encoding/json"
  "io"
  "net/http"
  "bytes"
)

type td struct {
  Uid int `json:"uid"`
  Id int `json:"id"`
  Title string `json:"title"`
  Completed bool `json:"completed"`
}

func main() {
  todo := td{Uid: 1, Id: 2, Title: "text", Completed: true}
  jr, err := json.Marshal(todo)
  if (err != nil) {
    fmt.Println(err)
    return
  }
  url := "https://jsonplaceholder.typicode.com/posts/1"
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jr))
  if (err != nil) {
    fmt.Println(err)
    return
  }
  req.Header.Set("Content-type", "application/json; charset=UTF-8")
  cli := &http.Client{}
  resp, err := cli.Do(req)
  if (err != nil) {
    fmt.Println(err)
    return
  }
  defer resp.Body.Close()
  bb, err := io.ReadAll(resp.Body)
  if (err != nil) {
    fmt.Println(err)
    return
  }
  bs := string(bb)
  fmt.Printf("%s\n\n", bs)
  var tds td
  err = json.Unmarshal(bb, &tds)
  if (err != nil) {
    fmt.Println(err)
    return
  }
  fmt.Printf("%s\n\n", tds)
  fmt.Printf(resp.Status)
}
