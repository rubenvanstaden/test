SHELL := /bin/bash

UNIT_TEST_PATH=./...

tidy:
	GOPROXY="direct" go mod tidy -compat=1.17
	gofmt -l -s -w .

test.unit:
	go test -count=1 -run=Unit $(UNIT_TEST_PATH)

test.unit.debug:
	go test -count=1 -run=Unit $(UNIT_TEST_PATH) -v
