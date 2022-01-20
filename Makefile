GO := go

all: build

## --------------------------------------
##@ General
## --------------------------------------

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
##@ Development
## --------------------------------------

fmt: ## Run go fmt against code.
	$(GO) fmt ./...

vet: ## Run go vet against code.
	$(GO) vet ./...

mod: ## Run go mod tidy.
	$(GO) mod tidy

build: mod fmt vet ## Build all binaries.
	$(GO) build -o bin/kubemulti