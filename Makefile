.PHONY: build

default: build

BUILD_TIME=`date +%FT%T%z`
GIT_REVISION=`git rev-parse --short HEAD`
GIT_BRANCH=`git rev-parse --symbolic-full-name --abbrev-ref HEAD`
GIT_DIRTY=`git diff-index --quiet HEAD -- || echo "✗-"`

LDFLAGS=-ldflags "-X main.buildTime=${BUILD_TIME} -X main.gitRevision=${GIT_DIRTY}${GIT_REVISION} -X main.gitBranch=${GIT_BRANCH}"

build:
	go build -o cmd/codegen/codegen ${LDFLAGS} cmd/codegen/main.go

install:
	cd cmd/codegen && go install ${LDFLAGS} && cd ../..
	