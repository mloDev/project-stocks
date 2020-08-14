IMAGE ?= project-stocks
REGISTRY ?= eu.gcr.io
VERSION ?= latest
DIRS += $(shell find */* -maxdepth 0 -name Makefile -exec dirname "{}" \;)
TMPGOPATH := $(shell mktemp -d)

precommit: ensure
	@echo "ready to commit"

ensure:
	GO111MODULE=on go mod verify
	GO111MODULE=on go mod vendor

deploy:
	mv vendor $(TMPGOPATH)/src
	GOPATH=$(TMPGOPATH) gcloud app deploy
	mv $(TMPGOPATH)/src vendor

test:
	go test -p=1 -cover -race $(shell go list ./... | grep -v /vendor/)

format:
	GO111MODULE=off go get github.com/seibert-media/goimports-reviser
	find . -type f -name '*.go' -not -path './vendor/*' -exec gofmt -w "{}" +

check: lint vet errcheck

vet:
	@go vet $(shell go list ./... | grep -v /vendor/)

lint:
	GO111MODULE=off go get golang.org/x/lint/golint
	@golint -min_confidence 1 $(shell go list ./... | grep -v /vendor/)

errcheck:
	GO111MODULE=off go get github.com/kisielk/errcheck
	@errcheck -ignore '(Close|Write|Fprint)' $(shell go list ./... | grep -v /vendor/)