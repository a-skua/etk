.PHONY: fmt test run

fmt:
	@go fmt ./...

test:
	@go test -v -cover ./...

run:
	@go run ./example
