package main

import (
  "bufio"
  "fmt"
  "net"
  "strings"
)

func main() {
  ln, err := net.Listen("tcp", ":8080")
  if err != nil {
    fmt.Println("Could not create listener due to error: ", err)
  }
  defer ln.Close()
  fmt.Println("Listening on port:", 8080)
  for {
    conn, err := ln.Accept()
    defer conn.Close()
    if err != nil {
      fmt.Println("count not accept connection, breaking out..")
      break
    }
    go handleConnection(conn)
  }
}

func handleConnection(conn net.Conn) {
  defer conn.Close()
  stringToRespond := request(conn)
  respond(conn, stringToRespond)
}

func request(conn net.Conn) string {
  scanner := bufio.NewScanner(conn)
  endpoint := ""
  i := 0
  for scanner.Scan() {
    line := scanner.Text()
    fmt.Println("connection said:", line)
    if i == 0 {
      endpoint = strings.Fields(line)[1]
    }
    if line == "" {
      break
    }
    i++
  }

  return endpoint
}

func respond(conn net.Conn, responseStr string) {
  body := `<!DOCTYPE html><html><head></head><title>This is test golang page.</title><body><h1>`+responseStr+`</h1></body></body></html>`
  fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
  fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
  fmt.Fprint(conn, "Content-Type: text/html\r\n")
  fmt.Fprint(conn, "\r\n")
  fmt.Fprint(conn, body)
}
