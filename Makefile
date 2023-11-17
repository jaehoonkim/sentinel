PACKAGE=github.com/jaehoonkim/synapse/pkg
VERSION=$(shell sed -n 's/VERSION=//p' properties.${target})
COMMIT=$(shell git rev-parse HEAD)
BUILD_DATE=$(shell date '+%Y-%m-%dT%H:%M:%S')
LDFLAGS=-X $(PACKAGE)/version.Version=$(VERSION) -X $(PACKAGE)/version.Commit=$(COMMIT) -X $(PACKAGE)/version.BuildDate=$(BUILD_DATE)

prep:
	go install github.com/swaggo/swag/cmd/swag@latest


swagger:
	cd pkg/manager/route;go generate

login:
	docker login ${register} -u ${user}

build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./bin/${target}/synapse-${target} ./cmd/${target}

image:
	docker build -t ${base}-${target}:latest -f Dockerfile.${target} .
	docker tag ${base}-${target}:latest ${base}-${target}:$(VERSION)

push:
	docker push ${base}-${target}:$(VERSION)
	docker push ${base}-${target}:latest

clean:
	rm ./bin/manager/synapse-manager
	rm ./bin/agent/synapse-agent
