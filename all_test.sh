#!/bin/bash

find . -name '*_test.go' | xargs go test -cover -v
