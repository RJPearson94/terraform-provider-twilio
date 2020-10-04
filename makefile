TEST?=$$(go list ./... | grep -v 'vendor' | grep -v 'examples')
PKG_NAME=twilio
TEST_COUNT?=1
ACCTEST_TIMEOUT?=60m
ACCTEST_PARALLELISM?=1

default: build

download: 
	@echo "==> Download dependencies"
	go mod vendor

build: fmt generate
	go install

test: fmt generate
	go test $(TESTARGS) -timeout=30s -parallel=4 $(TEST)

testacc: fmt
	# terraform is currently creating paths which are too long, so creating a temp directory to get around this issue
	mkdir -p temp
	TF_ACC_TEMP_DIR="$(CURDIR)/temp" TF_ACC=1 go test $(TEST) -v -count $(TEST_COUNT) -parallel $(ACCTEST_PARALLELISM) $(TESTARGS) -timeout $(ACCTEST_TIMEOUT)

fmt:
	@echo "==> Fixing source code with goimports (uses gofmt under the hood)..."
	goimports -w ./$(PKG_NAME) 
	gofmt -s -w ./$(PKG_NAME)

terrafmt:
	@echo "==> Format acceptance tests"
	@find twilio | egrep "_test.go" | sort | while read f; do terrafmt fmt -f $$f; done

terrafmt-docs:
	@echo "==> Format docs"
	@find docs | egrep ".md" | sort | while read f; do terrafmt fmt -f $$f; done

tools:
	@echo "==> installing required tooling..."
	GO111MODULE=off go get -u github.com/client9/misspell/cmd/misspell
	GO111MODULE=off go get -u github.com/bflad/tfproviderlint/cmd/tfproviderlint
	GO111MODULE=off go get -u github.com/bflad/tfproviderdocs
	GO111MODULE=off go get -u github.com/katbyte/terrafmt
	GO111MODULE=off go get -u github.com/boyter/scc
	GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports

reportcard:
	@echo "==> running go report card"
	goreportcard-cli

goreportcard-refresh:
	@echo "==> refresh goreportcard checks"
	curl -X POST -F "repo=github.com/RJPearson94/terraform-provider-twilio" https://goreportcard.com/checks

generate:
	go generate  ./...

.PHONY: download build test testacc fmt terrafmt terrafmt-docs tools generate reportcard goreportcard-refresh