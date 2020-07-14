.PHONY: generate
generate: generate-code generate-mocks tidy

.PHONY: generate-code
generate-code:
	rm -rf vendor_any
	go run generate.go
	./hack/post-generate.sh

.PHONY: generate-mocks
generate-mocks:
	go generate ./...

.PHONY: tidy
tidy:
	goimports -w .
	go mod tidy


.PHONY: mod-download
mod-download:
	go mod download

.PHONY: update-deps
update-deps: mod-download
	go get -v golang.org/x/tools/cmd/goimports@v0.0.0-20200423205358-59e73619c742
	go get -v github.com/gogo/protobuf/gogoproto@v1.3.1
	go get -v github.com/gogo/protobuf/protoc-gen-gogo@v1.3.1
	go get -v github.com/solo-io/protoc-gen-ext@v0.0.7
	go get -v github.com/golang/mock/mockgen@v1.4.3
	go get -v istio.io/tools/cmd/protoc-gen-jsonshim
