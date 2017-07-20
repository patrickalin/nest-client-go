mPWD := $(shell pwd)
GOPATH := $(shell go env GOPATH)
PROJECT := patrickalin/nest-client-go
ARTEFACT := nest-client-go

BUILD_LDFLAGS := '$(LDFLAGS)'

all: build

checks:
	@echo "checks --- Check deps"
	@(env bash $(PWD)/scripts/build/checkdeps.sh)
	@echo "Checking project is in GOPATH"
	@(env bash $(PWD)/scripts/build/checkgopath.sh ${PROJECT})
	@echo "checks ended"

getdeps: checks
	@echo "Installing golint" && go get -u github.com/golang/lint/golint
	@echo "Installing gocyclo" && go get -u github.com/fzipp/gocyclo
	@echo "Installing deadcode" && go get -u github.com/remyoudompheng/go-misc/deadcode
	@echo "Installing misspell" && go get -u github.com/client9/misspell/cmd/misspell
	@echo "Installing ineffassign" && go get -u github.com/gordonklaus/ineffassign
	@echo "Installing errcheck" && go get -u github.com/kisielk/errcheck
	@echo "Installing go-torch" && go get -u github.com/uber/go-torch 
	@echo "Installing bindata" && go get -u github.com/jteeuwen/go-bindata/

getFlame: 
	@echo "Installing FlameGraph" && git clone git@github.com:brendangregg/FlameGraph.git ${GOPATH}/src/github/FlameGraph

verifiers: getdeps vet fmt lint cyclo spelling deadcode errcheck

vet:
	@echo "Running $@ suspicious constructs"
	@go tool vet -atomic -bool -copylocks -nilfunc -printf -shadow -rangeloops -unreachable -unsafeptr -unusedresult *.go

fmt:
	@echo "Running $@ indentation and blanks for alignment"
	@gofmt -d *.go

lint:
	@echo "Running $@ style mistakes"
	@${GOPATH}/bin/golint -set_exit_status github.com/patrickalin/bloomsky-client-go/pkg...

ineffassign:
	@echo "Running $@"
	@${GOPATH}/bin/ineffassign .

cyclo:
	@echo "Running $@"
	@${GOPATH}/bin/gocyclo -over 100 http.go
	@${GOPATH}/bin/gocyclo -over 100 utils.go

deadcode:
	@${GOPATH}/bin/deadcode

spelling:
	@${GOPATH}/bin/misspell -error *.go

errcheck:
	@echo "Running $@"
	@${GOPATH}/bin/errcheck github.com/${PROJECT}

# Builds, runs the verifiers then runs the tests.
check: test
test: verifiers build
	@echo "Running all testing"
	@go test $(GOFLAGS) .

build:
	@echo "build"
	@go list -f '{{ .Name }}: {{ .Doc }}'
	@go generate
	@go build .

coverage: build
	@echo "Running all coverage"
	@./scripts/test/go-coverage.sh

clean:
	@echo "Cleaning up all the generated files"
	@find . -name '*.test' | xargs rm -fv
	@rm -rf build
	@rm -rf release
	@rm -rf coverage.*
	@rm -rf prof.cpu
	@rm -rf *.log
	@rm -rf torch.svg
	@rm -rf profile.*

documentation:
	@echo "listen on http://localhost:8081 ctrl+c stop"
	@(env bash $(PWD)/scripts/doc/doc.sh)

bench:
	@echo "Running $@"
	@go list -f '{{ .Name }}: {{ .Doc }}'
	@go test -bench . -cpuprofile prof.cpu

travisGihtub:
    @travis encrypt GITHUB_SECRET_TOKEN=$(GITHUB_SECRET_TOKEN) -a

torch: bench
	@echo "Running $@"
	@PATH=${PATH}:${GOPATH}/src/github/FlameGraph go-torch --binaryname ${ARTEFACT}.test -b prof.cpu
	@open torch.svg

pprofInteractif: bench
	@go tool pprof ${ARTEFACT}.test prof.cpu

pprofRaw: bench
	@go tool pprof -raw ${ARTEFACT}.test prof.cpu

tag: test
	@(env bash $(PWD)/scripts/git/tag.sh)