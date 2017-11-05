
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

.PHONY: run-json
run-json: build
	@MYVAR=myvar ./target/tmpl -tmpl=./test/template-test.tpl -data=./test/test-data.json

.PHONY: run-yaml
run-yaml: build
	@MYVAR=myvar ./target/tmpl -tmpl=./test/template-test.tpl -data=./test/test-data.yaml

.PHONY: clean
clean:
	rm -fr target
