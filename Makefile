#! /bin/env make

BINARY=msp-enroll

.PHONY: build test clean


build:
	go build -o ${BINARY} cmd/msp-enroll/main.go

test:
	go test -cover -covermode=atomic .

clean:
	@rm -f ${BINARY} 
