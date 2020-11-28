GOLANGCI_LINT ?= go run github.com/golangci/golangci-lint/cmd/golangci-lint

check: lint build test

build:
	@go build ./...

install:
	@go install

run: install
	@plantai serve

test:
	@go test -coverprofile=coverage.out ./...

cover:
	@go tool cover -func=coverage.out

lint:
	@$(GOLANGCI_LINT) run

clean:
	@rm -rf vendor coverage.out