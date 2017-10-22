
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

.PHONY: run
run: build
    @MYVAR=myvar ./target/tmpl -tmpl=./test/template-test.tpl
    @MYVAR=myvar ./target/tmpl -tmpl=./test/template-test.tpl -dest=./target/result.txt

.PHONY: clean
clean:
	rm -fr target
