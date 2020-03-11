export PKG_LIST := $(shell go list ./... | grep -v /vendor)

mod: ## get dependency packages
	go mod download
	go mod tidy

test: ## Run unit tests
	go test -v ${PKG_LIST}

race: ## Run race detector
	go test -race ${PKG_LIST}

coverage: ## https://penkovski.com/post/gitlab-golang-test-coverage/
	go test ${PKG_LIST} -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html;

lint: ## Lint package
	golangci-lint version
	golangci-lint run

.PHONY: mod test race coverage lint