#!/usr/bin/env bash
cd $GOPATH/src/github.com/viktor-br/links-manager-serve
java -jar ~/apps/swagger/swagger-codegen-cli.jar generate -i swagger.json -DpackageName=client -l go -o client/