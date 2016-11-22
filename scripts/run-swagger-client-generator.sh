#!/usr/bin/env bash
SWAGGER_DIR=$1
SWAGGER_PATH="$SWAGGER_DIR/swagger-codegen-cli.jar"

if [[ ! -d $SWAGGER_DIR ]]; then
    echo "Expect first parameter is directory with Swagger Codegen"
    exit 1
fi

if [[ ! -f $SWAGGER_PATH ]]; then
    echo "Cannot find Swagger Codegen jar: $SWAGGER_PATH"
    exit 1
fi

cd $GOPATH/src/github.com/viktor-br/links-manager-server
java -jar $SWAGGER_PATH generate -i swagger.json -DpackageName=client -l go -o client/