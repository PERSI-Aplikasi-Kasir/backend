dev:
	cls
	go run cmd/bin/main.go

migrate:
	cls
	go run cmd/bin/main.go --migrate

seed:
	cls
	go run cmd/bin/main.go --seed

logexposer:
	cls
	go run cmd/bin/main.go --logexposer


# Server
build:
	go build -o dist/ ./cmd/bin

start:
	./dist/bin --$(cmd)

deploy:
	go build -o dist/ ./cmd/bin
	./dist/bin
