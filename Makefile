pkgs = $(shell go list ./...)

.PHONY: build

# go build command
build:
	@go build -v -o mongo_helper cmd/*.go

# go run command
run:
	make build
	@./mongo_helper

test:
	@echo "RUN TESTING..."
	@go test -v -cover -race $(pkgs)