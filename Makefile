.PHONY: vendor tools

# Setup everything
setup: dep

download:
	go mod download

tidy:
	go mod tidy

vendor:
	go mod vendor

get:
	go get $(module)

get-all:
	go get -u all

# Setup go deps
dep: download tidy vendor

# Lint all the go code
lint:
	golangci-lint run --timeout 5m

# Fix the lint issues in the go code (if possible)
fix-lint:
	golangci-lint run --timeout 5m --fix

# Run all the specs
specs:
	go test -race -mod vendor -v -covermode=atomic -coverpkg=./... -coverprofile=profile.cov ./...

# Update go dep
update-dep: get tidy vendor

# Update all go dep
update-all-deps: get-all tidy vendor
