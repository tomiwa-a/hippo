.PHONY: ui run test tidy help

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## ui: run the frontend development server
ui:
	cd frontend && npm run dev

## run: run the Hippo crawler and engine
run:
	go run cmd/hippo/main.go

## test: run all tests, including ingestion verification
test:
	go test -v ./...

## tidy: clean up dependencies
tidy:
	go mod tidy
