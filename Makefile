.PHONY: clean
clean:
	rm -Rf bin/

.PHONY: unittests
unittests:
	go clean -testcache
	go test -v ./...

.PHONY: benchtests
benchtests:
	go clean -testcache
	go test -bench=. -v -benchmem -benchtime=2s ./tests/bench

.PHONY: build
build: clean
	go build -o bin/validgen .

.PHONY: endtoendtests
endtoendtests: build
	find tests/endtoend/ -name '*_validator.go' -exec rm \{} \;
	./bin/validgen tests/endtoend
	cd tests/endtoend; go run .
