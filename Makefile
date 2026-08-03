.PHONY: build install clean bump-version

BINARY_NAME := darkmatter
GO := go
GOFLAGS := -v
VERSION ?= $(shell cat .version 2>/dev/null || echo "latest")
BUILD_DIR := ./bin

help:
	clear;
	@echo "Darkmatter build help:"
	@echo "  make help          - Display this message"
	@echo "  make build         - Build the binary"
	@echo "  make install       - Install to GOBIN or GOPATH/bin"
	@echo "  make bump-version  - Automatic versioning, creates a new version"
	@echo "  make clean         - Remove build artifacts"

build:
	@echo ":: Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) -ldflags="-X main.version=$(VERSION)" -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/$(BINARY_NAME)

install:
	@echo ":: Installing $(BINARY_NAME)..."
	$(GO) install $(GOFLAGS) -ldflags="-X main.version=$(VERSION)" ./cmd/$(BINARY_NAME)

clean:
	@echo ":: Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)

bump-version:
	@read -p "Enter new version (current: $(shell cat .version 2>/dev/null || echo latest)): " NEW_VERSION; \
	VERSION_CLEAN=$$(echo "$$NEW_VERSION" | sed 's/^v//'); \
	FULL_TAG="v$$VERSION_CLEAN"; \
	echo "$$VERSION_CLEAN" > .version; \
	echo ":: Updating README.md..."; \
	sed "s|download/v[^/]*|download/$$FULL_TAG|g" README.md > README.tmp && mv README.tmp README.md; \
	sed "s|darkmatter-v.*-linux-amd64|darkmatter-$$FULL_TAG-linux-amd64|g" README.md > README.tmp && mv README.tmp README.md; \
	git add .version README.md; \
	git commit -S -m "chore: bump version to $$FULL_TAG"; \
	echo ":: Version bumped to $$FULL_TAG"; \
	echo ":: Don't forget to: git push && git tag $$FULL_TAG && git push --tags"
