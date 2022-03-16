MAKEFLAGS += --silent

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

DATABASE_URL=postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=${DATABASE_SSL_MODE}

init:
	docker run \
		--rm \
		-u 1100:1100 \
		-v ${PWD}:/opt/codeot \
		-w /opt/codeot \
		kjconroy/sqlc init

setup:
	go mod vendor

develop:
	go run main.go

dao:
	rm -rf ./app/dao/*
	docker run \
		--rm \
		-u 1000:1000 \
		-v ${PWD}:/opt/codeot \
		-w /opt/codeot \
		kjconroy/sqlc generate

test:
	go install gotest.tools/gotestsum@latest 1> /dev/null
	gotestsum --format testname --junitfile junit.xml -- -coverpkg=./app/... -coverprofile=cover.out ./app/... | grep -v "^coverage:"
	go tool cover -func cover.out | grep "total:"

mock:
	rm -rf mocks
	docker run \
		-it \
		--rm \
		-u 1000:1000 \
		-v ${PWD}:/service \
		-w /service \
		vektra/mockery --all --keeptree --dir app

db-migration:
	$(eval timestamp := $(shell date +%s))
	touch db/migrations/$(timestamp)_${name}.up.sql
	touch db/migrations/$(timestamp)_${name}.down.sql

db-rollback:
	docker run -v ${PWD}/db/migrations:/migrations \
		--rm -it --network host migrate/migrate \
		--path=/migrations/ \
		--database ${DATABASE_URL} down

db-migrate:
	docker run -v ${PWD}/db/migrations:/migrations \
		--rm -it --network host migrate/migrate \
		--path=/migrations/ \
		--database ${DATABASE_URL} up

db-create:
	createdb ${DATABASE_NAME}

db-drop:
	dropdb ${DATABASE_NAME}

db-seed:
	cp ./db/seeds.sql ./db/seeds.sql.bak
	envsubst < ./db/seeds.sql.bak > ./db/seeds.sql
	sed -i 's/COPY/\\copy/g' ./db/seeds.sql
	psql -a -f ./db/seeds.sql ${DATABASE_URL} 1> /dev/null
	mv ./db/seeds.sql.bak ./db/seeds.sql

db-clean: db-drop db-create db-migrate

db-setup: db-create db-migrate db-seed
