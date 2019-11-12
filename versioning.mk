MUTABLE_VERSION := latest
BUILD_TIME = `date +%FT%T%z`
VERSION := $(shell cat VERSION.txt 2>/dev/null)

GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY  = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")

ifdef VERSION
	DOCKER_VERSION = $(VERSION)
	BINARY_VERSION = $(VERSION)
endif

DOCKER_VERSION ?= git-${GIT_SHA}
BINARY_VERSION ?= ${GIT_TAG}

# Only set Version if building a tag or VERSION is set
ifneq ($(BINARY_VERSION),)
	LDFLAGS += -X main.Version=${BINARY_VERSION}
endif

# Clear the "unreleased" string in BuildMetadata
ifneq ($(GIT_TAG),)
	LDFLAGS += -X main.BuildMetadata=
endif
LDFLAGS += -X main.buildTime=${BUILD_TIME}
LDFLAGS += -X main.gitCommit=${GIT_COMMIT}
LDFLAGS += -X main.gitSha=${GIT_SHA}
LDFLAGS += -X main.gitTag=${GIT_TAG}
LDFLAGS += -X main.gitTreeState=${GIT_DIRTY}

IMAGE                := ${DOCKER_REGISTRY}/${IMAGE_PREFIX}/${SHORT_NAME}:${DOCKER_VERSION}
MUTABLE_IMAGE        := ${DOCKER_REGISTRY}/${IMAGE_PREFIX}/${SHORT_NAME}:${MUTABLE_VERSION}

info:
	 @echo "Build Time:        ${BUILD_TIME}"
	 @echo "Version:           ${VERSION}"
	 @echo "Git Tag:           ${GIT_TAG}"
	 @echo "Git Commit:        ${GIT_COMMIT}"
	 @echo "Git Tree State:    ${GIT_DIRTY}"
	 @echo "Docker Version:    ${DOCKER_VERSION}"
	 @echo "Registry:          ${DOCKER_REGISTRY}"
	 @echo "Immutable Image:   ${IMAGE}"
	 @echo "Mutable Image:     ${MUTABLE_IMAGE}"
