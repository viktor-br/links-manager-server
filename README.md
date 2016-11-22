# Links Manager
Service provides simple REST API ([API documentation](https://github.com/viktor-br/links-manager-server/blob/master/client/docs/DefaultApi.md)) to manage links and its tags (add, update, remove, search).

# Development

## Project structure
Try to implement [The Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html) (maybe the last who try).
```
/app 
    /controllers          functions, which generate HTTP response by calling interactors.
    /handlers             functions, which map resources to controllers.
    /links-manager-server main package of application.
    /log                  application logger interface.
    /mocks                mocks for unit tests.
/client                   API client code, auto-generated by Swagger Codegen
/core
    /config               configuration structures.
    /dao                  structures for a storage (and files are generated by ORM).
    /entities             core entities (user, session, link and so on).
    /implementation       implementation layer, which combines different sources.
    /interactors          main logic layer.
    /security             security urls.
/db                       db structure, test data and migrations.
/scripts                  scripts for test and development.
/tests                    end-to-end tests.
```

The application uses PostgreSQL as a storage and [reform](https://github.com/go-reform/reform) to interact with it.
You need to install reform:
```
$ go get -u gopkg.in/reform.v1
```
and run next command in case you need to regenerate DAO structure:
```
$ reform github.com/viktor-br/links-manager-server/core/dao
```

## API
Main source of API structure is in swagger.json

[Swagger Codegen](https://github.com/swagger-api/swagger-codegen) is used to regenerate REST API client (see /client folder).

```
$ scripts/run-swagger-client-generator.sh ~/apps/swagger
```

## Testing

Prerequisites: golang should be installed and project should be in $GOPATH; docker required.

Test layers, which doesn't need data storage (everything except implementation layer and end-to-end tests):
```
$ scripts/test-without-storage.sh
```

Test implementation layer, which requires storage (now it's postgresql):
```
$ scripts/run-test-storage.sh
$ scripts/test-implementation.sh
```

Run end-to-end tests, which use docker instances for application, storage and to run tests itself.
```
$ scripts/run-test-storage.sh
$ scripts/run-test-server.sh
$ scripts/test-acceptance.sh
```