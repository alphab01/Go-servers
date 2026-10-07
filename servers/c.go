package main

import (
  "fmt"
  "io"
  "log"
  "net/http"
  "net/url"
  "os"
  "time"
  "strconv"
)

var c int = 0

func h(w http.ResponseWriter, r *http.Request) {
  if (r.Method == "GET") {
    w.Write([]byte(strconv.Itoa(c)))
  } else if (r.Method == "POST") {
    r.ParseForm()
    s1 := r.Form.Get("count")
    c1, err := strconv.Atoi(s1)
    if (err == nil) {
      c += c1
    } else {
      w.WriteHeader(400)
      w.Write([]byte("это не число"))
    }
  }
}

func main() {
  http.HandleFunc("/count", h)
  _ = http.ListenAndServe(":3333", nil)
}
