BINDIR := $(CURDIR)/bin
BINNAME ?= todo
GOPATH := $(CURDIR)/vendor:$(CURDIR)
GODIR := $(CURDIR)/bin

	# @GOPATH=$(GOPATH) GOBIN=$(GODIR) go build -o $(BINDIR)/$(BINNAME) ./cmd/api/api.go
# ---
install:
	@make update-server

start-server:
	@make build-server
	@echo "	> Starting executable..."
	@$(BINDIR)/$(BINNAME)

update-server:
	@echo "	> Checking if there are any missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GODIR) go get $(get)

build-server:
	@echo "	> Building executable..."
	@GOPATH=$(GOPATH) GOBIN=$(GODIR) go build -o $(BINDIR)/$(BINNAME) -mod=mod cmd/$(BINNAME)/main.go

help:
	@echo "install: Install everything necessary to run the api"
	@echo "start-server: Build and start the api"
	@echo "build-server: Build the api into an executable"
	@echo "update-server: Get all the updates for the api"
