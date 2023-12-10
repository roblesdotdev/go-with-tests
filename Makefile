## test: test the application
.PHONY: test
test:
	go test ./... -v

## cover: test with coverage
.PHONY: cover
cover:
	go test ./... -cover