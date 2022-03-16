package main

import (
  "log"

  "github.com/fikrirnurhidayat/codeotd/app"
  "github.com/fikrirnurhidayat/codeotd/app/config"
)

func main() {
  port := config.GetPort()
  backend, err := app.New()

  if err != nil {
    log.Fatalf("Failed to initialize app: %v", err)
  }

  err = backend.ServeHTTP(port)

  if err != nil {
    log.Fatalf("Failed to serve http: %v", err)
  }
}
