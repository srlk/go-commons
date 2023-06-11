APP = experimenter
NS = experimenter

.PHONY: install-tools
install-tools:
	brew install golangci-lint

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run --fix -v ./...

