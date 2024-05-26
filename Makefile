migratecreate:
	migrate create -ext sql -dir internal/constants/query/schemas -tz "UTC" $(args)

migrateup:
	migrate -path internal/constants/query/schemas -database "postgres://yinebeb:yinebeb@localhost:5432/price_est?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/constants/query/schemas -database "postgres://yinebeb:yinebeb@localhost:5432/price_est?sslmode=disbale" -verbose down

sqlc:
	sqlc generate -f ./config/sqlc.yaml

swag:
	swag init -g initiator/initiator.go

test:
	go test -v -cover ./...

run:
	go run ./cmd/main.go

build:
	rm -f app
	go build -o app ./cmd/main.go

.PHONY: migratecreate migrateup migratedown sqlc swag test run build