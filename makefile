TEST?=$$(go list ./...)
PKG_NAME=twilio

default: build

download: 
	@echo "==> Download dependencies"
	go mod vendor

build: fmtcheck generate
	go install

test: fmtcheck generate
	go test $(TESTARGS) -timeout=30s -parallel=4 $(TEST)

testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 10m

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -w -s ./$(PKG_NAME)

fmtcheck:
	@echo "==> Checking that code complies with gofmt requirements..."
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

tools:
	@echo "==> installing required tooling..."
	go install github.com/client9/misspell/cmd/misspell
	go install github.com/katbyte/terrafmt

generate:
	go generate  ./...

.PHONY: download build test testacc fmt tools generate