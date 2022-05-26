DIST_LOCAL=dist/local
GOEXEC=go

lint:
	golangci-lint run .

build:
	$(GOEXEC) build -o $(DIST_LOCAL)/server cmd/http/server.go
