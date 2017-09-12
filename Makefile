.PHONY: build

ImageName=ymir	
ImageDestBase=hub.c.163.com/u2takey/ymir

PACKAGES = $(shell go list ./... | grep -v /vendor/ | grep -v /dashboard/)

ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
endif

BUILD_NUMBER=$(shell git rev-parse --short HEAD)


all: build_static

test:
	go test -cover $(PACKAGES)

build: build_static build_cross

build_static:
	mkdir -p make/release
	go build -o  make/release/ymir -ldflags '${EXTLDFLAGS}-X github.com/arlert/ymir/version.VersionDev=build.$(BUILD_NUMBER)' github.com/arlert/ymir/cmd

build_cross:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-X github.com/arlert/ymir/version.VersionDev=build.$(BUILD_NUMBER)' -o make/release/linux/amd64/ymir   github.com/arlert/ymir/cmd

build_tar:
	tar -cvzf make/release/linux/amd64/ymir.tar.gz   -C make/release/linux/amd64/ymir
	tar -cvzf make/release/darwin/amd64/ymir.tar.gz  -C make/release/darwin/amd64/ymir

build_docker:
	cd make && docker build -t $(ImageName) . && cd -

PushDest=$(ImageDestBase):v-$(shell date +'%y%m%d-%H%M%S')

publish_docker: build_docker
	docker tag $(ImageName) $(PushDest)
	docker push $(PushDest)

