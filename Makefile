all: dumptxs

dumptxs:
	go build -o ./cmd/dumptxs/dumptxs ./cmd/dumptxs/...

test:
	go test ./... -v -count=1

lint:
	golangci-lint run ./...
