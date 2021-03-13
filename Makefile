DIST_LOCAL=dist/local
GOEXEC=go

lint:
	$(GOEXEC) vet ./...

build:
	$(GOEXEC) build -o $(DIST_LOCAL)/cli cmd/cli/main.go
	$(GOEXEC) build -o $(DIST_LOCAL)/http cmd/http/main.go
