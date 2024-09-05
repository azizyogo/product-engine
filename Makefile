gorun:
	go run cmd/main.go

docker-build:
	docker build -t product-engine .

docker-run:
	docker compose up

create-migration:
	migrate create -ext json -dir ./migrations -seq $(name)

migrate-up:
	migrate -path ./migrations -database "mongodb://localhost:27017/product-engine" -verbose up

migrate-down:
	migrate -path ./migrations -database "mongodb://localhost:27017/product-engine" -verbose down