#
# Copyright (c) 2023. VMware, Inc. All rights reserved. VMware Confidential.
#

# install or update all build tools
tools = golang.org/x/lint/golint \
		golang.org/x/tools/cmd/goimports \
		github.com/vektra/mockery/v2

.PHONY: build-tools ${tools}
build-tools: ${tools}
${tools}: %: ; go install $*@latest

clean:
	rm -rf ./reports
	rm -rf ./bin

lint:
	golint ./...

vet:
	go vet ./...

sec:
	gosec ./...


.PHONY: mocks
mocks:
	@echo 'Deleting generated mocks...'
	@find ./generated/mocks -type f -delete
	@echo 'Running mockery...'
	@mockery


.PHONY: test
test:
	@echo 'Running unit tests...'
	@go test -cover $$(go list ./... | grep -v /generated/ | grep -v /cmd/)


.PHONY: convert
convert:
	@goverter gen -output-constraint '' -g 'output:package convert' -g skipCopySameType ./internal/converter	


# builds service executables for both eventmanager and API service
.PHONY: build
build:
	@go build -o ./bin/vss-cdm-eventmanager ./cmd/vss-cdm-eventmanager/main.go
	@go build -o ./bin/vss-cdm-service ./cmd/vss-cdm-service/main.go

.PHONY: run-eventmanager-local
run-eventmanager-local: build
	@INGRESS_SQS_NAME="vss-cdm-service-event-ingress" aws-vault exec dev-00 -- ./bin/vss-cdm-eventmanager

.PHONY: run-service-local
run-service-local: build
	@aws-vault exec dev-00 -- ./bin/vss-cdm-service

containers:
	docker build -f ./Dockerfile.service --build-arg PRIVATE_KEY="$(cat ~/.ssh/id_rsa)" .
	docker build -f ./Dockerfile.eventmanager --build-arg PRIVATE_KEY="$(cat ~/.ssh/id_rsa)" .

checks: lint vet sec test

# Upgrades the Go version on your local machine
.PHONY: upgrade
upgrade:
	@./scripts/upgrade.sh

# Upgrades the Go version in the files in the repo/project
.PHONY: upgrade-go
upgrade-go:
	@./scripts/upgrade.sh  Dockerfile.build Dockerfile.eventmanager Dockerfile.service Jenkinsfile-bellevue-ci

.PHONY: config
config:
	@config generate -i ./configs/config.yaml -o ./internal/generated/config -d ./CONFIG.md
	@goimports -w ./internal/generated/config

.PHONY: openapi openapi-update
openapi:
	@echo 'Deleting generated files...'
	@find ./internal/generated/api -type f -delete
	@echo 'Generating source files...'
	@openapi generate -clst -d=false -p api -i ./api/openapi.yaml -n CdmService -o ./internal/generated/api
	@goimports -w ./internal/generated/api
	@echo 'Generating public spec...'
	@openapi strip -m -i ./api/openapi.yaml -o ./web/openapi.yaml
openapi-update:
	@echo 'Downloading latest chss.yaml...'
	@git archive --format tar --remote git@gitlab.eng.vmware.com:securestate/go-sdk-v2.git HEAD:api chss.yaml | tar -x -C ./api

.PHONY: hub-api
hub-api:
	@echo 'Deleting generated files...'
	@find ./internal/generated/hub -type f -delete
	@echo 'Generating source files...'
	@openapi generate -l -p hub -i ./api/hub/openapi.yaml -o ./internal/generated/hub
	@goimports -w ./internal/generated/hub

# TODO: Event object for temporary use. This will be obsolete once we generate the event directly from the GraphQL schema.
.PHONY: endpointmanager-api
endpointmanager-api:
	@echo 'Deleting generated files...'
	@find ./internal/generated/endpointmanager -type f -delete
	@echo 'Generating source files...'
	@openapi generate -l -p endpoint_manager -i ./api/endpointmanager/openapi.yaml -o ./internal/generated/endpointmanager
	@goimports -w ./internal/generated/endpointmanager