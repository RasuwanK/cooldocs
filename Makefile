dev:
	go run ./cmd/server

test:
	go test ./...

db-up:
	docker compose up -d db 

db-down:
	docker compose down

db-reset:
	docker compose down -v
	docker compose up -d db 

migrate-up:
	goose -dir db/migrations up

migrate-down:
	goose -dir db/migrations down

docker-build:
	docker build -t cooldocs .

docker-up:
	docker compose up -d
