BINARY_NAME=ccwc

test:
	go test ./... -v

build:
	go build -o ${BINARY_NAME} main.go

clean:
	go clean
	rm ${BINARY_NAME}