#!/usr/bin/env bash

FILES=$(find . -name "*.go" | xargs -I % dirname % | sed 's/^.\///;s/[^.].*$/&\/*.go/;s/^\.$/*.go/' | sort -u)

echo "Running gofmt..."
res=$(gofmt -l ${FILES})
if [ -n "${res}" ]; then
  echo -e "format me please... \n${res}"
  exit 255
fi


echo "Running gometalinter..."
gometalinter -D gotype -D errcheck -D dupl ./... --deadline=120s


echo "Running tests..."
go test ./...
