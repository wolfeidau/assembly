default: test

deps:
	go get -v ./...

test: deps
	go test -v ./...

checkforgoreman: ; @which goreman > /dev/null || go get github.com/mattn/goreman

run: checkforgoreman
	goreman start

.PHONY: default deps test checkforgoreman run
