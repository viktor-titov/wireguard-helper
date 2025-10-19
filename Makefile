
.PHONY: build 

outdir != test -n "$$BUILD_OUTDIR" && echo $$BUILD_OUTDIR || echo ./build
project != basename $$PWD
sources := $(shell find . -name "*.go" -not -path "./vendor/*" | uniq)
version := $(shell git describe --tags)
email_sender := $(shell echo $$EMAIL_SENDER)
email_password := $(shell echo $$EMAIL_PASSWORD)


build:	$(sources)
	mkdir -p ${outdir}
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${outdir}/${project} \
	-ldflags "-X main.version=$(version) -X github.com/viktor-titov/wireguard-helper/internal/mail.sender=$(email_sender) -X 'github.com/viktor-titov/wireguard-helper/internal/mail.password=$(email_password)'" \
	./cmd
