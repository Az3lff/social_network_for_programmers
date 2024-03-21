.PHONY:
run:
	go run cmd/app/main.go

migrateup:
	migrate -path pkg/database/postgres/migrations -database "postgresql://postgres:marta2010@localhost:5432/socialNetwork" -verbose up