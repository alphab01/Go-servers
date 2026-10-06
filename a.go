package main

import (
  "fmt"
  "os"
  "net"
  "io"
)

func main() {
  req := "GET / HTTP/1.1.\n" + "Host: golang.org\n\n"
  c, err := net.Dial("tcp", "golang.org:80")
  if (err != nil) {
    fmt.Println(err)
    return
  }
  defer c.Close()
  _, err = c.Write([]byte(http.Request))
  if (err != nil) {
    fmt.Println(err)
    return
  }
  io.Copy(os.Stdout, c)
  fmt.Println("ok")
}
