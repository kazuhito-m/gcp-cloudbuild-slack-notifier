#!/bin/bash

for i in $(find . -name '*.go') ; do
  go fmt ${i}
done

result=0
for i in $(find . -name '*_test.go') ; do
  go test -cover -v ${i}

  r=$?
  if [ ${r} -ne 0 ]; then
    result=${r}
  fi  
done

exit ${result}