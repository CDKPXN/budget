.PHONY: build clean test docker

GO=CGO_ENABLED=0 GO111MODULE=on go

DOCKERS=docker_budget_api 
.PHONY: $(DOCKERS)

MICROSERVICES=budget-api
.PHONY: $(MICROSERVICES)

VERSION=$(shell cat ./VERSION)

# GOFLAGS=-ldflags "-X github.com/edgexfoundry/device-sdk-go.Version=$(VERSION)"

build: $(MICROSERVICES)

budget-api:
	$(GO) build -o $@ budget-api

clean:
	rm -f $(MICROSERVICES)

docker: $(DOCKERS)

docker_budget_api:
	docker build \
		-f Dockerfile \
		-t budget-api \
		.
test:
	GO111MODULE=on go test -v