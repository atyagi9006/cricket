test: clean
	@echo "mode: set" > coverage-all.out

	@go test -timeout 600s -tags unit -coverprofile=coverage.out -race ./... | \
		tee -a test-results.out || exit 1;\
		tail -n +2 coverage.out >> coverage-all.out || exit 1
	@go tool cover -html=coverage-all.out -o test-coverage.html

clean:
	rm -rf coverage* test-coverage.html test-results.out

run: clean
	@go run main.go

lint: 
	golint ./...

vet:
	go vet -v ./...

fmt:
	go fmt ./...

imports:
	goimports -l -w .

all: lint test  run