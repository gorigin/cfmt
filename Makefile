
all: vet build test

build:
	go build

fmt:
	go fmt

vet:
	go vet

test:
	go test -bench=. -benchmem

