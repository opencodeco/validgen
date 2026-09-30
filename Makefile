.PHONY: clean unittests build endtoendtests testgen setup lint

BIN_DIR=bin
BIN_NAME=validgen
VALIDGEN_BIN=$(BIN_DIR)/$(BIN_NAME)
ifdef VALIDGEN_BENCHMARKS_DIR
export VALIDGEN_BENCHMARKS_DIR := $(abspath $(VALIDGEN_BENCHMARKS_DIR))
endif
GOLANGCILINT_PATH=$(HOME)/bin
GOLANGCILINT_BIN=$(GOLANGCILINT_PATH)/golangci-lint

all: clean unittests build endtoendtests

clean:
	@echo "Cleaning"
	rm -Rf $(BIN_DIR)/

unittests:
	@echo "Running unit tests"
	go clean -testcache
	go test -v ./internal/... ./types/... ./testgen/

build: clean
	@echo "Building"
	go build -o $(VALIDGEN_BIN) .

testgen:
	@echo "Generating tests"
	cd testgen/ && rm -f generated_*.go && go run . && mv generated_endtoend_*tests.go ../tests/endtoend/ && mv generated_validation_*_test.go ../internal/codegenerator/ && mv generated_function_code_*_test.go ../internal/codegenerator/

endtoendtests: build
	@echo "Running endtoend tests"
	find tests/endtoend/ -name 'validator__.go' -exec rm \{} \;
	$(VALIDGEN_BIN) tests/endtoend
	cd tests/endtoend; go run .
	@echo "Running json unmarshal endtoend tests"
	find tests/jsonunmarshal/ -name 'validator__.go' -exec rm \{} \;
	$(VALIDGEN_BIN) -unmarshal-json tests/jsonunmarshal
	cd tests/jsonunmarshal; go run .

setup:
	@echo "Setting up"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(GOLANGCILINT_PATH) v2.5.0
	$(GOLANGCILINT_BIN) --version

lint:
	@echo "Linting"
	$(GOLANGCILINT_BIN) run --timeout=5m
