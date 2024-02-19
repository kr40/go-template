run: build
	@./bin/go-template

build: 
	@go build -o bin/go-template cmd/go-template/main.go

