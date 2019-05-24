PROJECT = nixug

VER=1.0.5.24
APP=nixug
SRC=github.com/ivost/nixug
REG=ivostoy
IMG=$(REG)/$(APP):$(VER)

SRCDIRS := ./...
#MOD="-mod=readonly"
MOD=
PKGS := $(shell go list -mod=readonly ./...)
HP := 0.0.0.0:8080

STYLE1=monokai
HTTP=http --style=$(STYLE1)
AUTH=--auth-type=jwt --auth=$(shell http http://localhost:8080/auth/nix/nix)

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
	go build -v cmd/service/nixug.go

install:
	go install -v cmd/service/nixug.go

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

######
docker:
	cd _deployment && docker build .  -t $(IMG)

push: docker
	docker push $(IMG)
pull:
	docker pull $(IMG)
drun:
	docker run --rm -d -p 8080:8080 $(IMG)
rund:
	docker run -it $(IMG)

######
h:
	$(HTTP) $(HP)/health

getg:
	$(HTTP) $(HP)/groups

getu:
	$(HTTP) $(HP)/users
