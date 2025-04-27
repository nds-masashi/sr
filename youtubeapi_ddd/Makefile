install:
	go install golang.org/x/tools/cmd/godoc@latest
	go install github.com/golang/mock/mockgen@latest
#	go install github.com/uber-go/mock/mockgen@latest

fmt:
	go fmt ./...

run:
	go run ./cmd/main.go

build:
	go build ./cmd/main.go

test:
	go test ./...

godoc:
	godoc