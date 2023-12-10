## test: test the application
.PHONY: test
test:
	go test ./...

## cover: test with coverage
.PHONY: cover
cover:
	go test ./... -cover