# social-go
migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users

migrate -path=./cmd/migrate/migrations -database="postgres://user:adminpassword@localhost/social?sslmode=disable" up

s