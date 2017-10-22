
.PHONY: default
default: build

.PHONY: build
build:
	@go fmt main.go
	@go fmt ./pkg/...
	@go build -o target/tmpl main.go

.PHONY: run
run: build
    @MYVAR=myvar ./target/tmpl -tmpl=./test/template-test.tpl
    @MYVAR=myvar ./target/tmpl -tmpl=./test/template-test.tpl -dest=./target/result.txt

.PHONY: clean
clean:
	rm -f target
