include .env

init:
	cp .env.example .env
	git config --local core.hooksPath githooks

start:
	docker compose up -d --build

migrate-up:
	docker run --rm -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:12000/${DB_DATABASE}?sslmode=disable up

migrate-down:
	docker run --rm -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:12000/${DB_DATABASE}?sslmode=disable down -all

migrate-create:
	docker run --rm -v ./migrations:/migrations migrate/migrate create -dir=/migrations -ext sql -seq create_$(name)_table

migrate-fix:
	docker run --rm -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:12000/${DB_DATABASE}?sslmode=disable force $(version)

swag:
	docker run --rm -v $(shell pwd):/app -w /app ghcr.io/swaggo/swag:latest init -g cmd/main.go

swag-fmt:
	docker run --rm -v $(shell pwd):/app -w /app ghcr.io/swaggo/swag:latest fmt