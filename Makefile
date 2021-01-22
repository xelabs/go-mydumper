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
	echo 'mode: atomic' > coverage.txt
	go test -covermode=atomic -coverprofile=coverage.out $(COVPKGS)
	go tool cover -html=coverage.out

.PHONY: all get build clean fmt test coverage
