dev:
	cls
	go run cmd/bin/main.go

migrate:
	cls
	go run cmd/bin/main.go --migrate

seed:
	cls
	go run cmd/bin/main.go --seed

logger:
	cls
	go run cmd/bin/main.go --logger

build:
	go build -o dist/ ./cmd/bin

start-win:
	dist/bin.exe

start-linux:
	./dist/bin

deploy:
	go build -o dist/bin ./cmd/bin
	./dist/bin
