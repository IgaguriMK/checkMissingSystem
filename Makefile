.PHONY: all
all: build test

.PHONY: build
build: checkMissing

.PHONY: checkMissing
checkMissing:
	go build checkMissing.go

.PHONY: test
test:
	make -C checker test
