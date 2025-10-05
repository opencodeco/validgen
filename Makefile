.PHONY: clean unittests benchtests build endtoendtests cmpbenchtests

BIN_DIR=bin
BIN_NAME=validgen
VALIDGEN_BIN=$(BIN_DIR)/$(BIN_NAME)
BENCH_TIME=5s
GOLANGCILINT_PATH=$(HOME)/bin
GOLANGCILINT_BIN=$(GOLANGCILINT_PATH)/golangci-lint

all: clean unittests build endtoendtests benchtests cmpbenchtests

clean:
	@echo "Cleaning"
	rm -Rf $(BIN_DIR)/

unittests:
	@echo "Running unit tests"
	go clean -testcache
	go test -v ./...

benchtests: build
	@echo "Running bench tests"
	find tests/bench/ -name '*_validator.go' -exec rm \{} \;
	$(VALIDGEN_BIN) tests/bench
	go clean -testcache
	go test -bench=. -v -benchmem -benchtime=$(BENCH_TIME) ./tests/bench

build: clean
	@echo "Building"
	go build -o $(VALIDGEN_BIN) .

endtoendtests: build
	@echo "Running endtoend tests"
	find tests/endtoend/ -name 'validator__.go' -exec rm \{} \;
	cd tests/endtoend/generate_tests/; rm -f numeric.go; go run generate_numeric_tests_main.go; mv numeric.go ..
	$(VALIDGEN_BIN) tests/endtoend
	cd tests/endtoend; go run .

cmpbenchtests: build
	@echo "Running cmp bench tests"
	rm -f tests/cmpbenchtests/generated_tests/*
	cd tests/cmpbenchtests; go run .
	$(VALIDGEN_BIN) tests/cmpbenchtests/generated_tests
	go clean -testcache
	go test -bench=. -v -benchmem -benchtime=$(BENCH_TIME) ./tests/cmpbenchtests/generated_tests

setup:
	@echo "Setting up"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(GOLANGCILINT_PATH) v2.5.0
	$(GOLANGCILINT_BIN) --version

lint:
	@echo "Linting"
	$(GOLANGCILINT_BIN) run --timeout=5m
