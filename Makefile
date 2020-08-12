OUT_DIR=bin

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
 
build: ## build app for dev
	@go build -mod=mod -o $(OUT_DIR)/devenv

full: clean mkdir shards-install build ## perform full build
	# done

build-release: clean mkdir deps-install ## perform release build
	@go build -ldflags "-s -w" -mod=mod -o $(OUT_DIR)/devenv

clean: ## clean compilation destination folder
	rm $(OUT_DIR)/* || true

mkdir: ## create compilation destination
	@mkdir -p $(OUT_DIR)

deps-install: ## install dependencies
	@go mod vendor
