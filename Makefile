
###
# Testing
##

.PHONY: create-mocks
create-mocks: ## Create the mocks for unit testing
	cd internal && go generate ./...

.PHONY: go-test
go-test: create-mocks ## Run the unit tests
	go test ./... -coverprofile=coverage.out

.PHONY: go-test-coverage
go-test-coverage: go-test ## Run the unit tests and display code coverage
	go tool cover -html=coverage.out