TEST?=$$(go list ./... | grep -v 'vendor' | grep -v 'examples')
PKG_NAME=twilio
TEST_COUNT?=1
ACCTEST_TIMEOUT?=60m
ACCTEST_PARALLELISM?=1
EXAMPLES?=$$(find . -type f -name "main.tf" -prune -exec dirname {} \;)

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

terraform-fmt:
	@echo "==> Format terraform examples"
	@find examples | egrep ".tf$$" | sort | while read f; do terraform fmt $$f; done

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

clean-examples:
	@echo "==> Cleaning examples"
	@find ./examples/* -type d -name ".terraform" -prune -exec rm -rf {} \;
	@find ./examples/* -type f -name "terraform.tfstate" -prune -exec rm -rf {} \;
	@find ./examples/* -type f -name "terraform.tfstate.backup" -prune -exec rm -rf {} \;
	@find ./examples/* -type f -name ".terraform.lock.hcl" -prune -exec rm -rf {} \;

validate-example:
	@echo "==> Validating example $(EXAMPLE)"
	terraform -chdir=$(EXAMPLE) init
	terraform -chdir=$(EXAMPLE) validate

validate-all-examples:
	@echo "==> Validating examples"
	make clean-examples
	for example in $(EXAMPLES); do \
		make validate-example EXAMPLE=$$example; \
	done

.PHONY: download build test testacc fmt terraform-fmt terrafmt terrafmt-docs tools generate reportcard goreportcard-refresh validate-example validate-all-examples clean-examples