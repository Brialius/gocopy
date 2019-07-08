VERSION := $(shell git describe --tags --dirty --always --match=v* || echo v0)
BUILD := $(shell git rev-parse --short HEAD)
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

.PHONY: setup
setup: ## Install all the build and lint dependencies
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint \
	golang.org/x/tools/cmd/cover \
	github.com/kisielk/errcheck \
    mvdan.cc/unindent \
    github.com/fzipp/gocyclo

.PHONY: test
test: ## Run all the tests
	go test -v ./...

.PHONY: lint
lint: ## Run all the linters
	golangci-lint run ./...
	gocyclo -over 15 ./
	errcheck
	unindent

.PHONY: ci
ci: lint test ## Run all the tests and code checks

.PHONY: build
build: ## Build a version
	go build $(LDFLAGS) -v

.PHONY: install
install: ## Build a version
	go install $(LDFLAGS) -v

.PHONY: clean
clean: ## Remove temporary files
	go clean

.PHONY: version
version:
	@echo $(VERSION)-$(BUILD)

.DEFAULT_GOAL := build
