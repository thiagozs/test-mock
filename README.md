# Golang mockgen test implementation

Stack of test

1. Golang 1.12.8
2. Mockgen
3. Gmock

Install packages

## Gomock

`go get -u -v github.com/golang/mock/gomock`

## Mockgen

`go install -v github.com/golang/mock/mockgen`

## Generate mocks by mockgen command

`mockgen -package mocks -destination mocks/redis.go -source dbs/redis.go DispatchServices, DispatchRepository`

## Build with

1. Redis

## Versioning and license

We use SemVer for versioning. You can see the versions available by checking the tags on this repository.

For more details about our license model, please take a look at the LICENSE.md file.

2019, thiagozs.