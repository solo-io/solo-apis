.PHONY: generate
generate: generate-code generate-mocks tidy

.PHONY: generate-code
generate-code:
	rm -rf vendor_any
	go run generate.go

.PHONY: generate-mocks
generate-mocks:
	go generate ./...

.PHONY: tidy
tidy:
	goimports -w pkg
	go mod tidy