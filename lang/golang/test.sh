#!/usr/bin/env bash

set -e;

dir="$(cd `dirname "$0"` && pwd)"
export GOPATH="$dir"

cd "$dir"

export CGOFILES=main
go install main