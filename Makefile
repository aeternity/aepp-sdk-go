GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES_UNIT = $(shell go list ./...  | grep -v /vendor/ | grep -v integration_test)
GIT_DESCR = $(shell git describe --tags --always)
APP=aecli
OUTPUTFOLDER = dist
ARCH = amd64

default: build

build: $(GOFILES)
	@echo build binary to $(OUTPUTFOLDER)
	CGO_ENABLED=0 go build -ldflags "-X main.Version=$(GIT_DESCR)" -o $(OUTPUTFOLDER)/$(APP) .
	@echo copy resources
	cp -r README.md LICENSE $(OUTPUTFOLDER)
	@echo done

build-release:
	@echo build binary to $(OUTPUTFOLDER)
	$(eval OS=darwin)
	GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=0 go build -ldflags "-X main.Version=$(GIT_DESCR)" -o $(OUTPUTFOLDER)/$(OS)/$(APP) .
	cp -r README.md LICENSE CHANGELOG.md $(OUTPUTFOLDER)/$(OS)
	zip -r $(OUTPUTFOLDER)/$(OS).zip $(OUTPUTFOLDER)/$(OS)
	$(eval OS=windows)
	GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=0 go build -ldflags "-X main.Version=$(GIT_DESCR)" -o $(OUTPUTFOLDER)/$(OS)/$(APP).exe .
	cp -r README.md LICENSE CHANGELOG.md $(OUTPUTFOLDER)/$(OS)
	zip -r $(OUTPUTFOLDER)/$(OS).zip $(OUTPUTFOLDER)/$(OS)
	$(eval OS=linux)
	GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=0 go build -ldflags "-X main.Version=$(GIT_DESCR)" -o $(OUTPUTFOLDER)/$(OS)/$(APP) .
	cp -r README.md LICENSE CHANGELOG.md $(OUTPUTFOLDER)/$(OS)
	zip -r $(OUTPUTFOLDER)/$(OS).zip $(OUTPUTFOLDER)/$(OS)
	@echo done

test: test-unit test-integration

test-unit:
	go test -v $(GOPACKAGES_UNIT)

test-integration:
	go test -v ./integration_test

bench:
	go test -bench -v ./...

lint:
	golint -set_exit_status ./...

install: build
	@cp $(OUTPUTFOLDER)/$(APP) $(GOPATH)/bin
	@echo copied to GOPATH/bin

clean:
	@echo remove $(OUTPUTFOLDER) folder
	@rm -rf $(OUTPUTFOLDER)
	@echo done
