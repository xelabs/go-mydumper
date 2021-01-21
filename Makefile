export PATH := $(GOPATH)/bin:$(PATH)

all: build test

build:
	@echo "--> Building..."
	go build -v -o bin/mydumper cmd/mydumper
	go build -v -o bin/myloader cmd/myloader
	@chmod 755 bin/*

clean:
	@echo "--> Cleaning..."
	@go clean
	@rm -f bin/*

fmt:
	gofumpt -w -s -d .
	go vet ./...

test:
	@echo "--> Testing..."
	@$(MAKE) testcommon

testcommon:
	go test -race -v common

# code coverage
COVPKGS =	common

coverage:
	@$(MAKE) get
	@$(MAKE) gettest
	go build -v -o bin/gotestcover \
	src/github.com/pierrre/gotestcover/*.go;
	bin/gotestcover -coverprofile=coverage.out -v $(COVPKGS)
	go tool cover -html=coverage.out

.PHONY: all get build clean fmt test coverage
