build:
	go build -o ginkgo-linter ./cmd/ginkgolinter

build-for-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/ginkgo-linter-amd64.exe ./cmd/ginkgolinter

build-for-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/ginkgo-linter-amd64-darwin ./cmd/ginkgolinter

build-for-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/ginkgo-linter-amd64-linux ./cmd/ginkgolinter
	GOOS=linux GOARCH=386 go build -o bin/ginkgo-linter-386-linux ./cmd/ginkgolinter

build-all: build build-for-linux build-for-mac build-for-windows
