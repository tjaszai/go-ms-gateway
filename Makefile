build:
	cd ./cmd && wire && cd ..
	swag init --dir ./cmd,./internal --output ./docs
	go build -o release/app cmd/main.go cmd/wire_gen.go

start:
	./release/app

run: build
	make start

cls:
	rm -rf ./release/

fmt:
	go fmt ./...

analyze:
	go vet ./...
