DIST_LOCAL=dist/local
GOEXEC=go

lint:
	golangci-lint run .

bench:
	go test -bench=. -count=5
test:
	go test -count=1 ./...
build: lint
	$(GOEXEC) build -o $(DIST_LOCAL)/server cmd/http/server.go

run: build
	./dist/local/server
