
.PHONY: default
default: build

.PHONY: test
test:
	@go vet .
	@go vet ./pkg/...
	@go fmt .
	@go fmt ./pkg/...
	@go test .
	@go test ./pkg/...

.PHONY: build
build: test
	@go build -o target/tmpl main.go

.PHONY: install
install: build
	@cp target/tmpl $(GOPATH)/bin

.PHONY: clean
clean:
	rm -fr target

.PHONY: test-all
test-all:
	docker run --rm -v $(shell pwd):/project -w /project alpine:3.5 ./test/test_it.sh
