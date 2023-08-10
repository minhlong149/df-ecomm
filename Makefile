.PHONY: run db

db:
	docker-compose up -d db

run:
	go run cmd/server/main.go
