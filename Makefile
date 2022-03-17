SHELL = /usr/bin/env bash

PKGNAME := scrawl
BINNAME := scrawl

.PHONY: %

all: test

test:
	@go test -v -cover ./...

build: clean
	@mkdir -p build/$$(go env GOOS)-$$(go env GOARCH)
	@go build -o build/$$(go env GOOS)-$$(go env GOARCH)/$(BINNAME)

xbuild: clean
	@for target in \
		darwin-amd64 \
		linux-amd64 \
		linux-arm \
		linux-arm64 \
		windows-amd64 \
	; \
	do \
		GOOS=$$(echo "$${target}" | cut -d'-' -f1) ; \
		GOARCH=$$(echo "$${target}" | cut -d'-' -f2) ; \
		outdir=build/"$${GOOS}-$${GOARCH}" ; \
		mkdir -p "$${outdir}" ; \
		printf "Building for %s-%s into build/ ...\n" "$${GOOS}" "$${GOARCH}" ; \
		GOOS="$${GOOS}" GOARCH="$${GOARCH}" go build -o "$${outdir}"/$(BINNAME) ./... ; \
	done

package: xbuild
	@mkdir -p dist
	@cd build || exit 1; \
	for built in * ; do \
		printf 'Packaging for %s into dist/ ...\n' "$${built}" ; \
		cd $${built} && tar -czf ../../dist/$(PKGNAME)_$${built}.tar.gz * && cd - >/dev/null ; \
	done

clean:
	@rm -rf \
		/tmp/$(PKGNAME)-tests \
		build/ \
		dist/
