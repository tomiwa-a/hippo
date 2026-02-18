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
	go run ./cmd/hippo start

## status: check engine status
status:
	go run ./cmd/hippo status

## stop: stop the engine
stop:
	go run ./cmd/hippo stop

## query: search the index (usage: make query Q="search term")
query:
	go run ./cmd/hippo query "$(Q)"

## test: run all tests, including ingestion verification
test:
	go test -v ./...

## tidy: clean up dependencies
tidy:
	go mod tidy
