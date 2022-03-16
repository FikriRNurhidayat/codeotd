CREATE TABLE IF NOT EXISTS test_cases (
  id uuid DEFAULT uuid_generate_v4(),
  name VARCHAR(255) NOT NULL,
  hidden BOOLEAN NOT NULL,
  input TEXT NOT NULL,
  output TEXT NOT NULL,
  challenge_id uuid NOT NULL REFERENCES challenges(id),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NULL,
  PRIMARY KEY (id)
);
