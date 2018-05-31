ROOT_DIR=$(shell pwd)
DOCKER_IMAGE=puppetlabs/bootstrap
OS_TYPE=$(shell echo `uname`| tr '[A-Z]' '[a-z]')
TMP_DIR=$$PWD/tmp

lint:
	golint $$(go list ./... | grep -v /vendor/)

test:
	go test -cover $$(go list ././... | grep -v /vendor/) 

format:
	go fmt $$(go list ./... | grep -v /vendor/) 

vet:
	go vet $$(go list ./... | grep -v /vendor/)

image:
	docker build -t ${DOCKER_IMAGE} .

shell:
ifeq ($(OS_TYPE), linux)
	if [ -d $(TMP_DIR) ]; then sudo rm -rf  $(TMP_DIR); fi 
	if [ ! -d $(TMP_DIR) ]; then mkdir $(TMP_DIR); fi
	docker build -t ${DOCKER_IMAGE}-shell -f hack/Dockerfile.shell . 
	docker run -it --rm --privileged -v $(TMP_DIR):/var/lib/containerd ${DOCKER_IMAGE}-shell /bin/bash || true
	sudo rm -rf $(TMP_DIR)  	
else
	vagrant up
	vagrant ssh -c 'cd /vagrant/hack && sudo ./entrypoint.sh'
	vagrant ssh 	
endif

clean:
ifeq ($(OS_TYPE), linux)
	docker rmi ${DOCKER_IMAGE} || true
	docker rmi ${DOCKER_IMAGE}-shell || true
else 
	vagrant destroy || true
	rm -rf $$PWD/bin
	rm -rf $$PWD/src
	rm -rf $$PWD/pkg
endif	
