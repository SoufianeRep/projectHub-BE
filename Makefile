postgres:
	docker run --name projecthub-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it projecthub-db createdb --username=root --owner=root project_hub

dropdb:
	docker exec -it projecthub-db dropdb project_hub


# migrateup:
# 	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/project_hub?sslmode=disable" -verbose up

# migratedown:
# 	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/project_hub?sslmode=disable" -verbose down

server:
	go run main.go

.PHONY: prostgres createdb dropdb server
