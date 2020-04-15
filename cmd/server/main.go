package main

import (
  "log"
  "github.com/thinkmassive/proglog/internal/server"
)

func main() {
  srv := server.NewHTTPServer(":8080")
  log.Fatal(srv.ListenAndServe())
}
