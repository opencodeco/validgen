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

.PHONY: end2endtests
end2endtests: build
	find tests/end2end/ -name '*_validator.go' -exec rm \{} \;
	./bin/validgen tests/end2end
	cd tests/end2end; go run .
