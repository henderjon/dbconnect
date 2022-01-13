.PHONY: mod
mod:
	go mod tidy
	go mod vendor

.PHONY: check
check: mod
	golint
	goimports -w ./
	gofmt -w ./
	go vet
