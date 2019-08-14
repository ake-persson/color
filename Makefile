all:    test

prereq:
	go get -u golang.org/x/lint/golint
  
format:
	gofmt -w .

test: format
	golint -set_exit_status .
	go vet .

.PHONY: prereq format test
