# -----------------------------------------------
#  Definitions
# -----------------------------------------------
GO 		 := go

BINDIR 	 := $(CURDIR)/bin
BINNAME  := golang-api-template
MAINPATH := cmd/$(BINNAME)/main.go

# -----------------------------------------------
#  Commands
# -----------------------------------------------
## help: Print this message
help:
	@fgrep -h '##' $(MAKEFILE_LIST) | fgrep -v fgrep | column -t -s ':' | sed -e 's/## //'

## build: Create the binary 
build:
	@echo "Building executable in $(BINDIR)/$(BINNAME)"
	@$(GO) build -o $(BINDIR)/$(BINNAME) -mod=mod $(MAINPATH)

## run: Run the binary
run:
	@$(BINDIR)/$(BINNAME)

## vendor: Download the vendored dependencies 
vendor:
	@$(GO) mod vendor

## lint: Lint the project 
lint:
	@golangci-lint run

## test: Test the project
#test:
	#$(GO) test 
