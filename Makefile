all: test

.PHONY: test

setup:
	go get golang.org/x/tools/cmd/goimports

test:
	go test -v -coverprofile=coverage.txt -covermode=count

gofmt:
	goimports -w .
