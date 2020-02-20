DOCKER_REGISTRY   ?= registry.shizzle.de:443
IMAGE_PREFIX      ?= kanban
DEV_IMAGE         ?= golang:1.13
SHORT_NAME        ?= 
TARGETS           ?= linux/amd64
TARGET_OBJS       ?= linux-amd64.tar.gz linux-amd64.tar.gz.sha256
DIST_DIRS         = find * -type d -exec

# go option
GO        ?= go
TAGS      :=
TESTS     := .
TESTFLAGS :=
LDFLAGS   := -w -s
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin
BINARIES  := codegen
BINARIES_DOCKER = $(patsubst %,docker-%, $(BINARIES))
BINARIES_PUSH = $(patsubst %,push-%, $(BINARIES))
BINARIES_BUMP = $(patsubst %,bump-%, $(BINARIES))

# tools
CP := cp -u -v

# Required for globs to work correctly
SHELL=/usr/bin/env bash

.PHONY: all
all: build

.PHONY: build
build: $(BINARIES)

.PHONY: $(BINARIES)
$(BINARIES):
	echo "building $@"
	$(eval LDF_LOCAL := ${LDFLAGS})
	$(eval VERSION_FILE := cmd/$@/VERSION.txt)
	$(eval VERSION := $(shell cat ${VERSION_FILE}))
	$(eval LDF_LOCAL += -X main.version=${VERSION})
	fileb0x b0x.toml
	$(GO) build $(GOFLAGS) -tags '$(TAGS)' -o ./dist/$@  -ldflags "$(LDF_LOCAL)" ./cmd/$@
	echo "done building $@ -> dist/$@"

.PHONY: $(BINARIES_DOCKER)
$(BINARIES_DOCKER): docker-%: % check-docker
	docker build --rm -f Dockerfile.$(patsubst docker-%,%, $@) -t $(patsubst docker-%, %, $@) .

.PHONY: $(BINARIES_PUSH)
$(BINARIES_PUSH): check-docker
	$(eval VERSION_FILE := cmd/$(patsubst push-%,%, $@)/VERSION.txt)
	$(eval VERSION := $(shell cat ${VERSION_FILE}))
	$(eval DOCKER_VERSION := $(shell cat ${VERSION_FILE}))
	docker tag $(patsubst push-%,%, $@):latest ${DOCKER_REGISTRY}/$(patsubst push-%,%, $@):${MUTABLE_VERSION}
	docker tag $(patsubst push-%,%, $@):latest ${DOCKER_REGISTRY}/$(patsubst push-%,%, $@):${DOCKER_VERSION}
	docker push ${DOCKER_REGISTRY}/$(patsubst push-%,%, $@):${DOCKER_VERSION}
	docker push ${DOCKER_REGISTRY}/$(patsubst push-%,%, $@):${MUTABLE_VERSION}

# github.com/jessfraz/junk/sembump download to gopath externally
.PHONY: bump-all
bump-all: $(BINARIES_BUMP)

.PHONY: $(BINARIES_BUMP)
BUMP := patch
$(BINARIES_BUMP):
	@echo "(${BUMP})ing, target: $(patsubst bump-%,%, $@)"

	$(eval VERSION_FILE := cmd/$(patsubst bump-%,%, $@)/VERSION.txt)
	$(eval VERSION := $(shell cat ${VERSION_FILE}))
	$(eval NEW_VERSION = $(shell sembump --kind $(BUMP) $(VERSION)))

	@echo "Bumping VERSION.txt from $(VERSION) to $(NEW_VERSION)"
	@echo $(NEW_VERSION) > ${VERSION_FILE}
#	@git add ${VERSION_FILE}
#	@git commit -vsam "Bump '$(patsubst bump-%,%, $@)' version to $(NEW_VERSION)"

.PHONY: check-docker
check-docker:
	@if [ -z $$(which docker) ]; then \
	  echo "Missing \`docker\` client which is required for development"; \
	  exit 2; \
	fi

.PHONY: test
test: build
test: TESTFLAGS += -race -v

.PHONY: docker-test
docker-test: docker-binary
docker-test: TESTFLAGS += -race -v

.PHONY: test-unit
test-unit:
	@echo
	@echo "==> Running unit tests <=="
	$(GO) test $(GOFLAGS) -run $(TESTS) $$(go list ./... | grep -v /vendor/) $(TESTFLAGS)
	
.PHONY: docker-test-unit
docker-test-unit: check-docker
	docker run \
		-v $(shell pwd):/go/src/k8s.io/helm \
		-w /go/src/k8s.io/helm \
		$(DEV_IMAGE) \
		bash -c "go test $(GOFLAGS) -run $(TESTS) $$(go list ./... | grep -v /vendor/) $(TESTFLAGS)"

.PHONY: generate
generate:
	$(GO) generate ./...

.PHONY: clean
clean:
	@rm -rf $(BINDIR) ./_dist
	rm -rf dist/*
	find static -type f -name '*.br' -delete
	find static -type f -name '*.gz' -delete

include versioning.mk

daoup:
	codegen -config codegen.json update

daogen:
	codegen -config codegen.json gen

install:
	cd dist && cp codegen ~/gocode/bin/ && cd ..