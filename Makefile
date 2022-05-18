BINARY_NAME=numary
PKG=./...
NUMARY_STORAGE_DRIVER="postgres"
NUMARY_STORAGE_POSTGRES_CONN_STRING="postgresql://ledger:ledger@127.0.0.1/ledger"
ENABLED_LINTERS=gofmt,gci
FAILFAST=-failfast
TIMEOUT=10m
RUN=".*"

all: lint test

build:
	go build -o $(BINARY_NAME)

install: build
	cp $(BINARY_NAME) $(shell go env GOPATH)/bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s -- -b $(shell go env GOPATH)/bin latest

lint:
	golangci-lint run -v -E $(ENABLED_LINTERS) --fix $(PKG)

test: test-sqlite test-postgres

test-sqlite:
	go test -tags json1 -v $(FAILFAST) -coverpkg $(PKG) -coverprofile coverage.out -covermode atomic -run $(RUN) -timeout $(TIMEOUT) -race $(PKG) \
		| sed ''/PASS/s//$(shell printf "\033[32mPASS\033[0m")/'' \
		| sed ''/FAIL/s//$(shell printf "\033[31mFAIL\033[0m")/'' \
		| sed ''/RUN/s//$(shell printf "\033[34mRUN\033[0m")/''

test-postgres: postgres
	NUMARY_STORAGE_DRIVER=$(NUMARY_STORAGE_DRIVER) \
	NUMARY_STORAGE_POSTGRES_CONN_STRING=$(NUMARY_STORAGE_POSTGRES_CONN_STRING) \
	go test -tags json1 -v $(FAILFAST) -coverpkg $(PKG) -coverprofile coverage.out -covermode atomic -run $(RUN) -timeout $(TIMEOUT) -race $(PKG) \
		| sed ''/PASS/s//$(shell printf "\033[32mPASS\033[0m")/'' \
		| sed ''/FAIL/s//$(shell printf "\033[31mFAIL\033[0m")/'' \
		| sed ''/RUN/s//$(shell printf "\033[34mRUN\033[0m")/''

postgres:
	docker-compose up -d postgres

bench:
	go test -tags json1 -bench=. -run=^a $(PKG)

clean:
	docker-compose down -v
	go clean
	rm -f $(BINARY_NAME)
