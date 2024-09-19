dev:
	cls
	go run cmd/bin/main.go --$(cmd)

# Server
build:
	go build -o dist/ ./cmd/bin

start:
	./dist/bin --$(cmd)

deploy:
	go build -o dist/ ./cmd/bin
	./dist/bin
