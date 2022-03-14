package postgres

import (
  // Why is Go like this :"(
  // Utilizing SIDE EFFECT
  _ "github.com/lib/pq"
  "database/sql"
)

func ConnectDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

  return db, nil
}
