BINDIR := $(CURDIR)/bin
BINNAME ?= todo
GOPATH := $(CURDIR)/vendor:$(CURDIR)
GODIR := $(CURDIR)/bin

	# @GOPATH=$(GOPATH) GOBIN=$(GODIR) go build -o $(BINDIR)/$(BINNAME) ./cmd/api/api.go
# ---
start-server:
	@make build-server
	@echo "	> Starting executable..."
	@$(BINDIR)/$(BINNAME)

build-server:
	@echo "	> Building executable..."
	@GOPATH=$(GOPATH) GOBIN=$(GODIR) go build -o $(BINDIR)/$(BINNAME) -mod=mod cmd/$(BINNAME)/main.go

help:
	@echo "start-server: Build and start the api"
	@echo "build-server: Build the api into an executable"
