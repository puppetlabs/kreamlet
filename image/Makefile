KUBE_RUNTIME ?= cri-containerd
KUBE_NETWORK ?= weave

KUBE_NETWORK_WEAVE ?= v2.2.1

KUBE_FORMATS ?= aws gcp vhd iso-bios

KUBE_FORMAT_ARGS := $(patsubst %,-format %,$(KUBE_FORMATS))

KUBE_BASENAME ?= kube-

.PHONY: all masters nodes
all: masters nodes

masters: yml/kube.yml yml/$(KUBE_RUNTIME).yml yml/$(KUBE_RUNTIME)-master.yml yml/$(KUBE_NETWORK).yml $(KUBE_EXTRA_YML)
	for kube_format in $(KUBE_FORMATS); do \
		linuxkit $(LINUXKIT_ARGS) build $(LINUXKIT_BUILD_ARGS) -name $(KUBE_BASENAME)master -format $$kube_format $^; \
	done

nodes: yml/kube.yml yml/$(KUBE_RUNTIME).yml yml/$(KUBE_NETWORK).yml $(KUBE_EXTRA_YML)
	for kube_format in $(KUBE_FORMATS); do \
	    linuxkit $(LINUXKIT_ARGS) build $(LINUXKIT_BUILD_ARGS) -name $(KUBE_BASENAME)node -format $$kube_format $^; \
	done

yml/weave.yml: kube-weave.yaml

kube-weave.yaml:
	curl -L -o $@ https://cloud.weave.works/k8s/v1.8/net?v=$(KUBE_NETWORK_WEAVE)

.PHONY: update-hashes
update-hashes:
	set -e ; for tag in $$(linuxkit pkg show-tag pkg/kubelet) \
	           $$(linuxkit pkg show-tag pkg/cri-containerd) \
	           $$(linuxkit pkg show-tag pkg/kubernetes-docker-image-cache-common) \
	           $$(linuxkit pkg show-tag pkg/kubernetes-docker-image-cache-control-plane) ; do \
	    image=$${tag%:*} ; \
	    sed -E -i.bak -e "s,$$image:[[:xdigit:]]{40}(-dirty)?,$$tag,g" yml/*.yml ; \
	done

.PHONY: clean
clean:
	rm -f -r \
	  kube-*-kernel kube-*-cmdline kube-*-state kube-*-initrd.img *img.tar.gz *.vhd *.raw *.iso \
	  kube-weave.yaml

.PHONY: refresh-image-caches
refresh-image-caches:
	./scripts/mk-image-cache-lst common > pkg/kubernetes-docker-image-cache-common/images.lst
	./scripts/mk-image-cache-lst control-plane > pkg/kubernetes-docker-image-cache-control-plane/images.lst
