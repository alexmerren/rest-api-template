# -----------------------------------------------
#  Definitions
# -----------------------------------------------
GO := go

BIN_DIR := $(CURDIR)/bin
PKG_DIR := $(CURDIR)/pkg

BINNAME  := golang-api-template
MAINPATH := cmd/$(BINNAME)/main.go

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
	@$(GO) build -o $(BIN_DIR)/$(BINNAME) -mod=mod $(MAINPATH)

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

## test: Test the project
.PHONY: test
test:
	@$(GO) test -coverpkg=$(PKG_DIR)/... \
		-coverprofile=coverage.out  \
		$(PKG_DIR)/... 
	@$(GO) tool cover -func=coverage.out
	@$(GO) tool cover -html=coverage.out -o coverage.html 
