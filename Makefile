# -----------------------------------------------
#  Definitions
# -----------------------------------------------
BINDIR := $(CURDIR)/bin
BINNAME ?= golang-api-template
GOPATH := $(CURDIR)/vendor:$(CURDIR)
GODIR := $(CURDIR)/bin
CMDPATH := cmd/$(BINNAME)/main.go

# -----------------------------------------------
#  Commands
# -----------------------------------------------
build:
	@echo "Building executable in $(BINDIR)/$(BINNAME)"
	@GOPATH=$(GOPATH) GOBIN=$(GODIR) go build -o $(BINDIR)/$(BINNAME) -mod=vendor $(CMDPATH)

run:
	@$(BINDIR)/$(BINNAME)

vendor:
	@go mod vendor

lint:
	@golangci-lint run

all: vendor lint build run 

help:
	@echo "build: Build the api into an executable."
	@echo "lint: Lint the project."
	@echo "vendor: Download all the dependencies for the project."
	@echo "run: Run the api executable."
	@echo "all: Build and run the api."
