TEST?=$$(go list ./...)
PKG_NAME=twilio

default: build

download: 
	@echo "==> Download dependencies"
	go mod vendor

build: fmt generate
	go install

test: fmt generate
	go test $(TESTARGS) -timeout=30s -parallel=4 $(TEST)

testacc: fmt
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 10m

fmt:
	@echo "==> Fixing source code with goimports (uses gofmt under the hood)..."
	goimports -w ./$(PKG_NAME)

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

generate:
	go generate  ./...

.PHONY: download build test testacc fmt terrafmt terrafmt-docs tools generate