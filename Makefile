


.PHONY: default
default: build

.PHONY: build
build:
	go build -o target/tmpl main.go

.PHONY: clean
clean:
	rm -f target
