#!/usr/bin/env bash

# Run tests

cd $GOPATH/src/github.com/viktor-br/links-manager-server
docker run -it --rm --link lm-test-main-storage \
    -v $(pwd)/db:/tmp/lm-db postgres \
  psql -h lm-test-main-storage -d test -U postgres -f /tmp/lm-db/test-data.sql

cd $GOPATH
docker run -it --rm --name lm-acceptance-tests \
--link lm-test-server \
-e LMS_TEST_SERVER_BASE_URL="http://lm-test-server:8080/api/v1" \
-v $(pwd)/src:/go/src golang \
go test github.com/viktor-br/links-manager-server/tests