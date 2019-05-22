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

help:
	@echo "==="
	@echo "build commands"
	@echo "==="

	@echo "build    - build ./nixug binary"
	@echo "install  - install nixug binary (assuming GOPATH is in PATH)"
	@echo "run      - build and run app"
	@echo "test     - unit tests"
	@echo "testv    - unit tests with output"
	@echo "testr    - unit tests with race detection"
	@echo "testi    - integration tests"

	@echo check    - test test-race vet fmt
	@echo scheck   - static analysis
	@echo pedantic - check unparam errcheck

	@echo h     - health check
	@echo getu  - get users
	@echo getg  - get groups
	@echo ---

run:
	go run cmd/service/nixug.go

build:
	go build cmd/service/nixug.go

install:
	go install cmd/service/nixug.go

test:
	go test $(MOD) ./...

testv:
	go test $(MOD) ./... -v

testr:
	go test -race $(MOD) ./...

testi: build
	@echo Integration tests
	go test internal/test/integration_test.go  -tags=integration -v

vet:
	go vet ./...

fmt:
	go fmt ./...

check: test testr vet fmt

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

