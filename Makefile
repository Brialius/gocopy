VERSION ?= $(shell git describe --tags --dirty --always --match=v* || echo v0)
BUILD := $(shell git rev-parse --short HEAD)
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"
MODFLAGS=-mod vendor
BUILDFLAGS=$(MODFLAGS) $(LDFLAGS)
PROJECTNAME=gocopy
GOEXE := $(shell go env GOEXE)
BIN=bin/$(PROJECTNAME)$(GOEXE)

.PHONY: setup
setup: mod-refresh ## Install all the build and lint dependencies
	go install github.com/golangci/golangci-lint/cmd/golangci-lint \
	golang.org/x/tools/cmd/cover

.PHONY: test
test: ## Run all the tests
	go test -v $(BUILDFLAGS) ./...

.PHONY: lint
lint: ## Run all the linters
	golangci-lint run --enable-all --disable gochecknoinits --disable gochecknoglobals --disable goimports \
	--out-format=tab --tests=false ./...

.PHONY: ci
ci: setup lint test build ## Run all the tests and code checks

.PHONY: build
build: ## Build a version
	go build $(BUILDFLAGS) -o $(BIN)

.PHONY: install
install: ## Install a binary
	go install $(BUILDFLAGS)

.PHONY: clean
clean: ## Remove temporary files
	go clean

.PHONY: mod-refresh
mod-refresh: clean ## Refresh modules
	go mod tidy -v
	go mod vendor

.PHONY: version
version:
	@echo $(VERSION)-$(BUILD)

.DEFAULT_GOAL := build
