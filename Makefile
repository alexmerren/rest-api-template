BINDIR := $(CURDIR)/bin
BINNAME ?= todo
GOFLAGS :=
TAGS :=
LDFLAGS :=

# ---

.PHONY: build
build: $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME): $(SRC)
	GO111MODULE=on go build $(GOFLAGS) -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o '$(BINDIR)'/$(BINNAME) ./cmd/api
