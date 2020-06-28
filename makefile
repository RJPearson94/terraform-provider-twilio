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
	GO111MODULE=off go get -u github.com/client9/misspell/cmd/misspell
	GO111MODULE=off go get -u github.com/bflad/tfproviderlint/cmd/tfproviderlint
	GO111MODULE=off go get -u github.com/bflad/tfproviderdocs
	GO111MODULE=off go get -u github.com/katbyte/terrafmt
	GO111MODULE=off go get -u github.com/boyter/scc

generate:
	go generate  ./...

.PHONY: download build test testacc fmt tools generate