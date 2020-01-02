VERSION = $$(git describe --abbrev=0 --tags)
VERSION_DATE = $$(git log -1 --pretty='%ad' --date=format:'%Y-%m-%d' $(VERSION))
COMMIT_REV = $$(git rev-list -n 1 $(VERSION))

all: build

version:
	@echo $(VERSION)

commit_rev:
	@echo $(COMMIT_REV)

start:
	go run nhl.go

deps/clean:
	go clean -modcache

debug:
	DEBUG=1 go run nhl.go

build:
	@go build -o bin/nhl nhl.go

# http://macappstore.org/upx
build/mac: clean/mac
	env GOARCH=amd64 go build -ldflags "-s -w" -o bin/macos/nhl && upx bin/macos/nhl

build/linux: clean/linux
	env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/linux/nhl && upx bin/linux/nhl

build/multiple: clean
	env GOARCH=amd64 go build -ldflags "-s -w" -o bin/nhl64 && upx bin/nhl64 && \
	env GOARCH=386 go build -ldflags "-s -w" -o bin/nhl32 && upx bin/nhl32

clean/mac:
	go clean && \
	rm -rf bin/mac

clean/linux:
	go clean && \
	rm -rf bin/linux

clean:
	go clean && \
	rm -rf bin/

test:
	go test ./...

nhl/test:
	go run nhl.go -test

nhl/version:
	go run nhl.go -version

nhl/clean:
	go run nhl.go -clean

nhl/reset:
	go run nhl.go -reset

git/repack:
	git reflog expire --expire=now --all
	git fsck --full --unreachable
	git repack -A -d
	git gc --aggressive --prune=now

release:
	rm -rf dist
	VERSION=$(VERSION) goreleaser