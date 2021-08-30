# -----------------------------------------------
#  Definitions
# -----------------------------------------------
BINDIR := $(CURDIR)/bin
BINNAME ?= todo
GOPATH := $(CURDIR)/vendor:$(CURDIR)
GODIR := $(CURDIR)/bin
CMDPATH := cmd/$(BINNAME)/main.go

# -----------------------------------------------
#  Commands
# -----------------------------------------------
build-server:
	@echo "	> Building executable..."
	@GOPATH=$(GOPATH) GOBIN=$(GODIR) go build -o $(BINDIR)/$(BINNAME) -mod=vendor $(CMDPATH)

start-server:
	@make build-server
	@echo "	> Starting executable..."
	@echo "	> Press Ctrl+C to stop the server..."
	@$(BINDIR)/$(BINNAME)

help:
	@echo "build-server: Build the api into an executable"
	@echo "start-server: Build and start the api"
