package config

import "fmt"

func GetPort() string {
  port := Getenv("PORT", "8000")
  return fmt.Sprintf(":%s", port)
}
