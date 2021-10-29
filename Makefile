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

all: build run

help:
	@echo "build: Build the api into an executable."
	@echo "run: Run the api executable."
	@echo "all: Build and run the api."
