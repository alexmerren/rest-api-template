# -----------------------------------------------
#  Definitions
# -----------------------------------------------
GO := go

BIN_DIR := $(CURDIR)/bin
INTERNAL_DIR := $(CURDIR)/internal

BINNAME  := rest-api-template
MAINPATH := cmd/$(BINNAME)/main.go

TEST_MODULES := $(shell $(GO) list $(INTERNAL_DIR)/...)

# -----------------------------------------------
#  Commands
# -----------------------------------------------
## help: Print this message
help:
	@fgrep -h '##' $(MAKEFILE_LIST) | fgrep -v fgrep | column -t -s ':' | sed -e 's/## //'

## build: Create the binary 
.PHONY: build
build:
	@echo "Building executable in $(BIN_DIR)/$(BINNAME)"
	@$(GO) build -o $(BIN_DIR)/$(BINNAME) -mod=vendor $(MAINPATH)

## run: Run the binary
.PHONY: run
run:
	@$(BIN_DIR)/$(BINNAME)

## vendor: Download the vendored dependencies 
.PHONY: vendor
vendor:
	@$(GO) mod tidy
	@$(GO) mod vendor

## lint: Lint the project 
.PHONY: lint
lint:
	@golangci-lint run

## test: Run the unit tests for the project 
.PHONY: test
test:
	@$(GO) test $(TEST_MODULES) -coverprofile=$(BIN_DIR)/coverage.out coverpkg=$(INTERNAL_DIR)/...
	@$(GO) tool cover -html=$(BIN_DIR)/coverage.out -o $(BIN_DIR)/test-coverage.html
	@$(GO) tool cover -func=$(BIN_DIR)/coverage.out \
		| awk '$$1 == "total:" {printf("Total coverage: %.2f%% of statements\n", $$3)}'

## mocks: Generate mocks for the project
.PHONY: mocks
mocks:
	@echo "Not implemented :("
