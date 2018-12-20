GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)
GIT_DESCR = $(shell git describe --tags --always) 
APP=aecli
# build output folder
OUTPUTFOLDER = dist
# docker image
DOCKER_IMAGE = aeternity/aepps-sdk-go
DOCKER_TAG = $(shell git describe --tags --always)
# build paramters
OS = linux
ARCH = amd64

.PHONY: list
list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs

default: build

workdir:
	mkdir -p dist

build: build-dist

build-dist: $(GOFILES)
	@echo build binary to $(OUTPUTFOLDER)
	GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=0 go build -ldflags "-X main.Version=$(GIT_DESCR)" -o $(OUTPUTFOLDER)/$(APP) .
	@echo copy resources
	cp -r README.md LICENSE $(OUTPUTFOLDER)
	@echo done

build-relase:
	@echo build binary to $(OUTPUTFOLDER)
	$(eval OS=darwin)
	GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=0 go build -ldflags "-X main.Version=$(GIT_DESCR)" -o $(OUTPUTFOLDER)/$(OS)/$(APP) .
	$(eval OS=windows)
	GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=0 go build -ldflags "-X main.Version=$(GIT_DESCR)" -o $(OUTPUTFOLDER)/$(OS)/$(APP).exe .
	$(eval OS=linux)
	GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=0 go build -ldflags "-X main.Version=$(GIT_DESCR)" -o $(OUTPUTFOLDER)/$(OS)/$(APP) .
	@echo copy resources
	cp -r README.md LICENSE CHANGELOG.md $(OUTPUTFOLDER)
	@echo done

test: test-all

test-all:
	@go test -v $(GOPACKAGES) -coverprofile .testCoverage.txt

bench: bench-all

bench-all:
	@go test -bench -v $(GOPACKAGES)

lint: lint-all

lint-all:
	@golint -set_exit_status $(GOPACKAGES)

install: build-dist
	@cp dist/aecli $(GOPATH)/bin
	@echo copied to GOPATH/bin

clean:
	@echo remove $(OUTPUTFOLDER) folder
	@rm -rf dist
	@echo done

docker: docker-build

docker-build: build-dist
	@echo copy resources
	@cp config/settings.docker.yaml $(OUTPUTFOLDER)/settings.yaml
	@echo build image
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) -f ./build/docker/Dockerfile .
	@echo done

docker-push: docker-build
	@echo push image
	docker tag $(DOCKER_IMAGE):$(DOCKER_TAG) $(DOCKER_IMAGE):latest
	docker push $(DOCKER_IMAGE)
	@echo done

docker-run: 
	@docker run -p 1804:1804 $(DOCKER_IMAGE) 

debug-start:
	@go run main.go -c config/settings.sample.yaml --debug start
