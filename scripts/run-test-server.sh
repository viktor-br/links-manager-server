#!/usr/bin/env bash
echo "Install LM server..."
go install github.com/viktor-br/links-manager-server/app/links-manager-server

cd $GOPATH

RUNNING=$(docker inspect -f {{.State.Running}} lm-test-server 2> /dev/null)

if [ "$RUNNING" = "" ]; then
    echo "Container doesn't exist. Run..."

    docker run -it -d --name lm-test-server \
    --link lm-test-main-storage \
    -e LMS_MAIN_STORAGE_CONNECTION=postgres://postgres@lm-test-main-storage:5432/test?sslmode=disable \
    -e LMS_MAIN_STORAGE_TYPE=postgres \
    -e LMS_SECRET=123 \
    -p 8080:8080 \
    -v $(pwd)/bin:/go/bin golang /go/bin/links-manager-server
fi

if [ "$RUNNING" = "false" ]; then
    echo "Container was stopped. Start..."
    docker start lm-test-server
fi