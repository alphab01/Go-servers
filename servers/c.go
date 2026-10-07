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

var c := 0

func h(w ResponseWriter, r *http.Request) {
  if (r.Method == "GET") {
    w.Write([]byte(string(c)))
  } else if (r.Method == "POST") {
    if (isOnlyDigits(r.URL.GET("count"))) {
      c += strconv.Atoi(r.URL.GET("count"))
    } else {
      w.Write([]byte("это не число"))
      w.WriteHeader(400)
    }
  }
}

func main() {

}
