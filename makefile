
main: cmd/main.go
	go build cmd/main.go

async: cmd/async.go
	go build cmd/async.go

all: main async

.DEFAULT_GOAL := all