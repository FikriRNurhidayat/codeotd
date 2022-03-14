package config

import "fmt"

func GetDatabaseConnectionString() string {
  return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
    Getenv("DATABASE_USER", "root"),
    Getenv("DATABASE_PASSWORD", ""),
    Getenv("DATABASE_HOST", "localhost"),
    Getenv("DATABASE_PORT", "5432"),
    Getenv("DATABASE_NAME", "bookstoresvc_development"),
    Getenv("DATABASE_SSL_MODE", "disable"),
  )
}
