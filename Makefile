build:
	@go build -o build/parser cmd/parser.go
run: build
	./build/parser $(file)
clean:
	@rm -rf build