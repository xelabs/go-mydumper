export PATH := $(GOPATH)/bin:$(PATH)

all: build test

build:
	@echo "--> Building..."
	go build -v -o bin/mydumper ./cmd/mydumper
	go build -v -o bin/myloader ./cmd/myloader
	@chmod 755 bin/*

clean:
	@echo "--> Cleaning..."
	@go clean
	@rm -f bin/*

fmt:
	go fmt ./...
	go vet ./...

test:
	@echo "--> Testing..."
	@$(MAKE) testcommon

testcommon:
	go test -race -v ./common

# code coverage
COVPKGS =	./common

coverage:
	sh -c "echo 'mode: atomic' > coverage.txt" &&
	go list ./... | xargs -n1 -I{} sh -c \
		'go test -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.txt' && \
		rm coverage.tmp
	go tool cover -html=coverage.out

.PHONY: all get build clean fmt test coverage
