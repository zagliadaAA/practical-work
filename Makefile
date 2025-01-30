#migrate-up:
#	goose -dir ./db/migrations postgres "host=localhost user=postgres dbname=medical_center sslmode=disable"

migrate-up:
	goose -dir ./db/migrations postgres "host=localhost port=5432 user=postgres dbname=medical_center sslmode=disable" up

migrate-create:
	goose -dir db/migrations create $(F) sql

lint:
	golangci-lint run ./...