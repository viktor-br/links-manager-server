#!/bin/sh
# Check if main storage container is running. If not start
RUNNING=$(docker inspect -f {{.State.Running}} lm-test-main-storage 2> /dev/null)

# Container doesn't exist. Create storage container and propagate the db structure.
if [ "$RUNNING" = "" ]; then
    docker run --name lm-test-main-storage -d -e POSTGRES_DB=test postgres

    # Sleep for a few seconds while postgresql will be available.
    # Need some smarter solution.
    sleep 5

    cd $GOPATH/src/github.com/viktor-br/links-manager-server

    docker run -it --rm --link lm-test-main-storage \
        -v $(pwd)/db:/tmp/lm-db postgres \
      psql -h lm-test-main-storage -d test -U postgres -f /tmp/lm-db/structure.sql
fi

# Check if container is stopped
if [ "$RUNNING" = "false" ]; then
    docker start lm-test-main-storage
fi

cd $GOPATH

docker run -it --rm --name lm-storage-test \
--link lm-test-main-storage \
-e LMS_TEST_MAIN_STORAGE_CONNECTION=postgres://postgres@lm-test-main-storage:5432/test?sslmode=disable \
-e LMS_TEST_MAIN_STORAGE_TYPE=postgres \
-e LMS_TEST_SECRET=123 \
-v $(pwd)/src:/go/src golang \
go test -cover github.com/viktor-br/links-manager-server/core/implementation