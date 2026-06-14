# Makefile

# Variables
GO = go
PODMAN = podman
BINARY_NAME = main
IMAGE_NAME = my-service
VERSION = $(shell more ./VERSION)
BUILD_DATE = $(shell date +%Y%m%d)
BUILD_TIMESTAMP = $(shell date +%Y%m%dT%H%M%SZ)
CKSUM_SCRIPT = import hashlib; print(hashlib.sha1(open('./$(BINARY_NAME)','rb').read()).hexdigest())
METRICS_SERVER_PORT ?= 9090

# Default target
all: clean podman-build

.PHONY: clean test

clean:
	rm ./$(BINARY_NAME)

# Run the service locally using go run
run:
	$(GO) run ./cmd/main

# Build a static, stripped binary for minimal final image
build:
	$(GO) mod download
	CGO_ENABLED=0 $(GO) build -ldflags="-s -w" -o ./$(BINARY_NAME) ./cmd/main

# Run unit tests
test:
	$(GO) test -v ./...

# Build the container image using podman
podman-build: build
	$(PODMAN) build \
		--compress=true \
		--layers=true \
		--format=oci \
		-t $(IMAGE_NAME):latest \
		-t $(IMAGE_NAME):$(BUILD_DATE) \
		-t $(IMAGE_NAME):$(BUILD_TIMESTAMP) \
		-t $(IMAGE_NAME):$(VERSION) \
		-t $(IMAGE_NAME):$(shell python3 -c "$(CKSUM_SCRIPT)") \
		--build-arg BUILD_DATE=$(BUILD_DATE) \
		--build-arg BUILD_TIMESTAMP=$(BUILD_TIMESTAMP) \
		--build-arg VERSION=$(VERSION) \
		--build-arg METRICS_SERVER_PORT=$(METRICS_SERVER_PORT) \
		.

# Run the container locally using podman
podman-run: podman-build
	$(PODMAN) run \
		--rm -it \
		-p$(METRICS_SERVER_PORT):$(METRICS_SERVER_PORT) \
		$(IMAGE_NAME):latest
