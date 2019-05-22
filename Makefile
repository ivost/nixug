PROJECT = nixug

STYLE1=monokai

SRCDIRS := ./...

#MOD="-mod=readonly"
MOD=

PKGS := $(shell go list -mod=readonly ./...)

HP := 0.0.0.0:8080
G := $(V1)/meta

HTTP=http --style=$(STYLE1)
AUTH=--auth-type=jwt --auth=$(shell http localhost:8484/v1/auth/nix/nix)

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
	@echo '==='
	@echo 'build commands'
	@echo '==='

	@echo run   - build and run app
	@echo test  - unit tests
	@echo testv - unit tests with output
	@echo testr - unit tests with race detection
	@echo check - test test-race vet fmt
	@echo scheck- static analysis
	@echo pedantic - check unparam errcheck

	@echo h     - health check
	@echo getu  - get users
	@echo getg  - get groups
	@echo ---

test: install
	go test $(MOD) ./...

testv: install
	go test $(MOD) ./... -v

testr: | test
	go test -race $(MOD) ./...

vet: | test
	go vet ./...

fmt:
	go fmt ./...

check: test testr vet fmt

install:
	@#go install $(MOD) -v -tags "" ./...
	go mod tidy
	go install $(MOD) -v  ./...

download:
	go mod download


scheck:
	@echo static analysis
	go install honnef.co/go/tools/cmd/staticcheck
	staticcheck \
		-checks all,-ST1003,-S1011 \
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

getg:
	$(HTTP) $(HP)/groups

