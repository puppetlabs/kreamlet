ROOT_DIR=$(shell pwd)
BUILD_DIR=./bin
DOCKER_IMAGE=puppet/kreamlet
OS_TYPE=$(shell echo `uname`| tr '[A-Z]' '[a-z]')

.PHONY: all build-dirs lint test format vet binary clean

build-dirs:
	if [ ! -d $(BUILD_DIR) ]; then mkdir $(BUILD_DIR); fi

lint:
	golint $$(go list ./... | grep -v /vendor/)

test:
	go test -cover $$(go list ././... | grep -v /vendor/)

format:
	go fmt $$(go list ./... | grep -v /vendor/)

vet:
	go vet $$(go list ./... | grep -v /vendor/)

binary: build-dirs
	if [ ! -d $$PWD/tmp/ ]; then mkdir $$PWD/tmp/; fi
	docker build -t ${DOCKER_IMAGE} -f ${ROOT_DIR}/hack/Dockerfile.build .
	docker run --name kream-build ${DOCKER_IMAGE}
	docker cp kream-build:/go/src/github.com/puppetlabs/kreamlet/bin/kream $$PWD/tmp/
	docker cp kream-build:/go/src/github.com/puppetlabs/kreamlet/bin/kream-darwin $$PWD/tmp/
	docker rm kream-build
	cp -rf $$PWD/tmp/* $$PWD/bin && rm -rf $$PWD/tmp/

shell:
	docker build -t ${DOCKER_IMAGE}-shell -f hack/Dockerfile.shell .
	docker run -it --rm ${DOCKER_IMAGE}-shell /bin/bash

bootstrap-dev:
	cd $$PWD/bootstrap && make image
	if [ -d $$PWD/image/kube-master-state ]; then rm -rf $$PWD/image/kube-master-state; fi 
ifeq ($(OS_TYPE), linux)	
	if [ -d $$PWD/image/kube-master-state ]; then rm -rf $$PWD/image/kube-master-state; fi
	cd $$PWD/image && KUBE_FORMATS=iso-bios make all
	cd $$PWD/image && linuxkit run -publish 6443:6443 -publish 50091:50091 --mem 4096 kube-master.iso  
else
	if [ -d $$PWD/image/kube-master-efi-state ]; then rm -rf $$PWD/image/kube-master-efi-state; fi
	cd $$PWD/image && KUBE_FORMATS=iso-efi make all
	cd $$PWD/image && linuxkit run --mem 4096 -publish 6443:6443 -publish 50091:50091 -iso --uefi kube-master-efi.iso
endif


clean: 
	docker rmi ${DOCKER_IMAGE} || true
	rm -rf  ${ROOT_DIR}/bin/*
	cd $$PWD/image && make clean
