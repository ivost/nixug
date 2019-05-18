PROJECT = nix_users

STYLE1=monokai

SRCDIRS := ./...

PKGS := $(shell go list -mod=readonly ./...)

HP := 0.0.0.0:8484
V1 := $(HP)/v1
META := $(V1)/meta

HTTP=http --style=$(STYLE1)
AUTH=--auth-type=jwt --auth=$(shell http localhost:8484/v1/auth/tap/T@p)

#S := 1556218054463
S := 1556315241286
E :=
L := 10

ID1=users/1
ID2=users/2
ID3=groups/1
ID4=groups/2

GIT_REF = $(shell git rev-parse --short=8 --verify HEAD)
VERSION ?= $(GIT_REF)

export GO111MODULE=on

run:
	go run cmd/service/main.go

help:
	@echo h     - health check
	@echo getu  - get user
	@echo getg  - get group
	@echo ---

test: install
	go test -mod=readonly ./...

testv: install
	go test -mod=readonly ./... -v

test-race: | test
	go test -race -mod=readonly ./...

vet: | test
	go vet ./...

fmt:
	go fmt ./...

check: test test-race vet gofmt

install:
	@#go install -mod=readonly -v -tags "" ./...
	go mod tidy
	go install -mod=readonly -v  ./...

download:
	go mod download


staticcheck:
	go install honnef.co/go/tools/cmd/staticcheck
	staticcheck \
		-checks all,-ST1003 \
		$(PKGS)

pedantic: check unparam errcheck

unparam:
	go install mvdan.cc/unparam
	unparam ./...

errcheck:
	go install github.com/kisielk/errcheck
	errcheck $(PKGS)

h:
	$(HTTP) $(HP)/health

getu:
	$(HTTP) $(META)/
