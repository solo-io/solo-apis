.PHONY: generate
generate: generate-code generate-mocks

.PHONY: generate-code
generate-code:
	go run generate.go

.PHONY: generate-mocks
generate-mocks:
	go generate ./...

.PHONY: tidy
tidy:
	goimports -w pkg
	go mod tidy