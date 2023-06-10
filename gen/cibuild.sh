#!/bin/sh
# Generates cross builds for all supported platforms.
#
# This script is used to build binaries for all supported platforms. Cgo is
# disabled to make sure binaries are statically linked.

set -o verbose

ret=0

# https://golang.org/doc/install/source#environment

CGO_ENABLED=0 GOOS=android   GOARCH=arm64    go build || ret=1
CGO_ENABLED=0 GOOS=darwin    GOARCH=amd64    go build || ret=1
CGO_ENABLED=0 GOOS=dragonfly GOARCH=amd64    go build || ret=1
CGO_ENABLED=0 GOOS=freebsd   GOARCH=386      go build || ret=1
CGO_ENABLED=0 GOOS=freebsd   GOARCH=amd64    go build || ret=1
CGO_ENABLED=0 GOOS=freebsd   GOARCH=arm      go build || ret=1
CGO_ENABLED=0 GOOS=illumos   GOARCH=amd64    go build || ret=1
CGO_ENABLED=0 GOOS=linux     GOARCH=386      go build || ret=1
CGO_ENABLED=0 GOOS=linux     GOARCH=amd64    go build || ret=1
CGO_ENABLED=0 GOOS=linux     GOARCH=arm      go build || ret=1
CGO_ENABLED=0 GOOS=linux     GOARCH=arm64    go build || ret=1
CGO_ENABLED=0 GOOS=linux     GOARCH=ppc64    go build || ret=1
CGO_ENABLED=0 GOOS=linux     GOARCH=ppc64le  go build || ret=1
CGO_ENABLED=0 GOOS=linux     GOARCH=mips     go build || ret=1
CGO_ENABLED=0 GOOS=linux     GOARCH=mipsle   go build || ret=1
CGO_ENABLED=0 GOOS=linux     GOARCH=mips64   go build || ret=1
CGO_ENABLED=0 GOOS=linux     GOARCH=mips64le go build || ret=1
CGO_ENABLED=0 GOOS=linux     GOARCH=s390x    go build || ret=1
CGO_ENABLED=0 GOOS=netbsd    GOARCH=386      go build || ret=1
CGO_ENABLED=0 GOOS=netbsd    GOARCH=amd64    go build || ret=1
CGO_ENABLED=0 GOOS=netbsd    GOARCH=arm      go build || ret=1
CGO_ENABLED=0 GOOS=openbsd   GOARCH=386      go build || ret=1
CGO_ENABLED=0 GOOS=openbsd   GOARCH=amd64    go build || ret=1
CGO_ENABLED=0 GOOS=openbsd   GOARCH=arm      go build || ret=1
CGO_ENABLED=0 GOOS=openbsd   GOARCH=arm64    go build || ret=1
CGO_ENABLED=0 GOOS=solaris   GOARCH=amd64    go build || ret=1

CGO_ENABLED=0 GOOS=windows   GOARCH=386      go build || ret=1
CGO_ENABLED=0 GOOS=windows   GOARCH=amd64    go build || ret=1

# not supported
# CGO_ENABLED=0 GOOS=aix       GOARCH=ppc64    go build || ret=1
# CGO_ENABLED=0 GOOS=android   GOARCH=386      go build || ret=1
# CGO_ENABLED=0 GOOS=android   GOARCH=amd64    go build || ret=1
# CGO_ENABLED=0 GOOS=android   GOARCH=arm      go build || ret=1
# CGO_ENABLED=0 GOOS=darwin    GOARCH=arm64    go build || ret=1
# CGO_ENABLED=0 GOOS=js        GOARCH=wasm     go build || ret=1
# CGO_ENABLED=0 GOOS=plan9     GOARCH=386      go build || ret=1
# CGO_ENABLED=0 GOOS=plan9     GOARCH=amd64    go build || ret=1
# CGO_ENABLED=0 GOOS=plan9     GOARCH=arm      go build || ret=1

exit $ret
