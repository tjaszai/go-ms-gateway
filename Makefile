build:
	go build -o release/app cmd/main.go

run: build
	./release/app

fmt:
	go fmt ./...

analyze:
	go vet ./...
