#!/bin/bash

for i in $(find . -name '*_test.go') ; do
  go test -cover -v ${i}
done
