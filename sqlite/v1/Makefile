BIN=
HEAD=$(shell git describe --dirty --long --tags 2> /dev/null  || git rev-parse --short HEAD)
TIMESTAMP=$(shell date '+%Y-%m-%dT%H:%M:%S %z %Z')
DEPLOYMENT_PATH=""

LDFLAGS=""

docker: # only needs to be be run once to build the docker env locally; use docker-cgo to iterate builds
	env GOOS=linux GOARCH=amd64 go build -ldflags $(LDFLAGS) -o $(BIN) ./cmd/
	docker build -t "$(BIN):$(HEAD)" .

# CGO
.PHONY: linux64-cgo
linux64-cgo:
	env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -ldflags $(LDFLAGS) -o $(BIN)-linux64-cgo-$(HEAD) ./cmd/

.PHONY: docker-cgo
docker-cgo: clean
	docker run --rm -v "$(PWD)":/usr/src/myapp -w /usr/src/myapp golang:1.12.4 make linux64-cgo

.PHONY: cgo
cgo: docker-cgo
	$(info cp "$(BIN)-linux64-cgo-$(HEAD)" "$(DEPLOYMENT_PATH)")
