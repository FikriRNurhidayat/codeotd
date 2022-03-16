-- Seeding tests
COPY challenges(id, title, description, body) FROM '${PWD}/db/data/challenges.csv' DELIMITER ',' CSV HEADER;

-- Seeding test_cases
COPY test_cases(id, name, hidden, input, output, challenge_id) FROM '${PWD}/db/data/test_cases.csv' DELIMITER ',' CSV HEADER;
