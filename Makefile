.PHONY: clean unittests benchtests build endtoendtests cmpbenchtests

BIN_DIR=bin
BIN_NAME=validgen
VALIDGEN_BIN=$(BIN_DIR)/$(BIN_NAME)
BENCH_TIME=5s

all: clean unittests build endtoendtests benchtests cmpbenchtests

clean:
	rm -Rf $(BIN_DIR)/

unittests:
	go clean -testcache
	go test -v ./...

benchtests: build
	find tests/bench/ -name '*_validator.go' -exec rm \{} \;
	$(VALIDGEN_BIN) tests/bench
	go clean -testcache
	go test -bench=. -v -benchmem -benchtime=$(BENCH_TIME) ./tests/bench

build: clean
	go build -o $(VALIDGEN_BIN) .

endtoendtests: build
	find tests/endtoend/ -name 'validator__.go' -exec rm \{} \;
	$(VALIDGEN_BIN) tests/endtoend
	cd tests/endtoend; go run .

cmpbenchtests: build
	rm -f tests/cmpbenchtests/generated_tests/*
	cd tests/cmpbenchtests; go run .
	$(VALIDGEN_BIN) tests/cmpbenchtests/generated_tests
	go clean -testcache
	go test -bench=. -v -benchmem -benchtime=$(BENCH_TIME) ./tests/cmpbenchtests/generated_tests
