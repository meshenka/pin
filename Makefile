DIST_LOCAL=dist/local
GOEXEC=go

lint:
	golangci-lint run .

build: lint
	$(GOEXEC) build -o $(DIST_LOCAL)/server cmd/http/server.go

run: build
	./dist/local/server
