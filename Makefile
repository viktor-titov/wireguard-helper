
.PHONY: build 

outdir != test -n "$$BUILD_OUTDIR" && echo $$BUILD_OUTDIR || echo ./build
project != basename $$PWD
sources := $(shell find . -name "*.go" -not -path "./vendor/*" | uniq)
version := $(shell git describe --tags)

build: $(sources)
	mkdir -p ${outdir}
	CGO_ENABLED=0 go build -o ${outdir}/${project} -ldflags "-X main.version=${version}" ./cmd
