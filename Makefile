clean:
	-@rm crawler 2>/dev/null || true
build: clean
	@go build -o crawler
test:
	@go test ./...
runMock: build
	@./crawler "https://wagslane.dev"
