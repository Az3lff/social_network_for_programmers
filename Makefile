.PHONY:
run:
	go run cmd/app/main.go

migrateup:
	migrate -path pkg/database/postgres/migrations -database "postgresql://az3lff:SergeyAndKirill@localhost:5432/social_network_for_programmers" -verbose up

migratedown:
	migrate -path pkg/database/postgres/migrations -database "postgresql://az3lff:SergeyAndKirill@localhost:5432/social_network_for_programmers" -verbose up