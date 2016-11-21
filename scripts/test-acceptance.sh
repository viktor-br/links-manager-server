#!/usr/bin/env bash

# Run tests
cd $GOPATH

docker run -it --rm --name lm-acceptance-tests \
--link lm-test-server \
-e LMS_TEST_SERVER_BASE_URL="http://lm-test-server:8080/api/v1" \
-v $(pwd)/src:/go/src golang \
go test github.com/viktor-br/links-manager-server/tests