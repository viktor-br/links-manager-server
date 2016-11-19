#!/bin/sh
go test -cover github.com/viktor-br/links-manager-server/app/controllers
go test -cover github.com/viktor-br/links-manager-server/app/handlers
go test -cover github.com/viktor-br/links-manager-server/app/log
go test -cover github.com/viktor-br/links-manager-server/core/config
go test -cover github.com/viktor-br/links-manager-server/core/dao
go test -cover github.com/viktor-br/links-manager-server/core/entities
go test -cover github.com/viktor-br/links-manager-server/core/interactors
go test -cover github.com/viktor-br/links-manager-server/core/security