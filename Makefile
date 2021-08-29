BINDIR := $(CURDIR)/bin
BINNAME ?= todo
GOPATH := $(CURDIR)/vendor:$(CURDIR)
GODIR := $(CURDIR)/bin

	# @GOPATH=$(GOPATH) GOBIN=$(GODIR) go build -o $(BINDIR)/$(BINNAME) ./cmd/api/api.go
# ---
install:
	go-get

go-get:
	@echo "	> Checking if there are any missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GODIR) go get $(get)

go-build:
	@echo "	> Building executable..."
	@GOPATH=$(GOPATH) GOBIN=$(GODIR) go build -o $(BINDIR)/$(BINNAME) -mod=mod main.go
