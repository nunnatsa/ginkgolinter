build:
	go build -o ginkgo-linter ./cmd

build-for-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/ginkgo-linter-amd64.exe ./cmd

build-for-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/ginkgo-linter-amd64-darwin ./cmd

build-for-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/ginkgo-linter-amd64-linux ./cmd
	GOOS=linux GOARCH=386 go build -o bin/ginkgo-linter-386-linux ./cmd

build-all: build build-for-linux build-for-mac build-for-windows
