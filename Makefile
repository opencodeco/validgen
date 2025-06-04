clean:
	rm -Rf bin/

test:
	go clean -testcache
	go test -v ./...

bench:
	go clean -testcache
	go test -bench=. -v -benchmem -benchtime=2s ./benchtests

build: clean
	go build -o bin/ ./validgen/...
