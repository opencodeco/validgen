BIN_DIR=bin
BIN_NAME=validgen
VALIDGEN_BIN=$(BIN_DIR)/$(BIN_NAME)
BENCH_TIME=5s

.PHONY: clean
clean:
	rm -Rf $(BIN_DIR)/

.PHONY: unittests
unittests:
	go clean -testcache
	go test -v ./...

.PHONY: benchtests
benchtests: build
	find tests/bench/ -name '*_validator.go' -exec rm \{} \;
	$(VALIDGEN_BIN) tests/bench
	go clean -testcache
	go test -bench=. -v -benchmem -benchtime=$(BENCH_TIME) ./tests/bench

.PHONY: build
build: clean
	go build -o $(VALIDGEN_BIN) .

.PHONY: endtoendtests
endtoendtests: build
	find tests/endtoend/ -name '*_validator.go' -exec rm \{} \;
	$(VALIDGEN_BIN) tests/endtoend
	cd tests/endtoend; go run .

.PHONY: cmpbenchtests
cmpbenchtests: build
	rm -f tests/cmpbenchtests/generated_tests/*
	cd tests/cmpbenchtests; go run .
	$(VALIDGEN_BIN) tests/cmpbenchtests/generated_tests
	go clean -testcache
	go test -bench=. -v -benchmem -benchtime=$(BENCH_TIME) ./tests/cmpbenchtests/generated_tests
