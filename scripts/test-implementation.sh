#!/usr/bin/env bash
# Run go generate for reform files
reform github.com/viktor-br/links-manager-server/core/dao

# Run tests
cd $GOPATH

docker run -it --rm --name lm-storage-test \
--link lm-test-main-storage \
-e LMS_TEST_MAIN_STORAGE_CONNECTION=postgres://postgres@lm-test-main-storage:5432/test?sslmode=disable \
-e LMS_TEST_MAIN_STORAGE_TYPE=postgres \
-e LMS_TEST_SECRET=123 \
-v $(pwd)/src:/go/src golang \
go test -cover github.com/viktor-br/links-manager-server/core/implementation