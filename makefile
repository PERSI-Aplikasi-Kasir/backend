dev:
	cls
	go run cmd/bin/main.go

migrate:
	cls
	go run cmd/bin/main.go --migrate

seed:
	cls
	go run cmd/bin/main.go --seed

build:
	go build -o dist/ ./cmd/bin

start-win:
	dist/bin.exe

start-linux:
	./dist

deploy:
	go build -o dist/bin ./cmd/bin
	./dist