#!/bin/bash

for i in $(find . -name '*.go') ; do
  go fmt ${i}
done

go test -cover -v ./...
