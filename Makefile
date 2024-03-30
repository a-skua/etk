.PHONY: fmt test run debug

fmt:
	@go fmt ./...

test:
	@go test -v -cover ./...

run:
	@go run ./example

debug:
	@go run -tags=ebitenginedebug ./example
