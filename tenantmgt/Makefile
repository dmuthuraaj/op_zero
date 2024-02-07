BINARY_NAME=tenantmgt

build:
	go build -o bin/${BINARY_NAME} cmd/main.go

run:
	go run cmd/main.go


compile:
	echo "Compiling for every OS and Platform..."
	GOOS=linux GOARCH=arm go build -o bin/${BINARY_NAME}-linux-arm cmd/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/${BINARY_NAME}-linux-arm64 cmd/main.go
	GOOS=freebsd GOARCH=386 go build -o bin/${BINARY_NAME}-freebsd-386 cmd/main.go

clean:
	go clean
	echo "Deleting binary files..."
	rm -rf bin

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

all: build run