SHELL := /bin/bash

UNIT_TEST_PATH=./...

fmt:
	gofmt -l -s -w .

tidy:
	GOPROXY="" go mod tidy -compat=1.17

test.unit:
	go test -count=1 -run=Unit $(UNIT_TEST_PATH)

test.unit.debug:
	go test -count=1 -run=Unit $(UNIT_TEST_PATH) -v
