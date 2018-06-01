# Kreamlet

Kreamlet is the service that orchestrates the orchestrator. Its job in life to to make sure the Kubernetes cluster is configured and running. For more information on this please see the docs folder.

## Development 

To develop kreamlet you will need the following things golang, containerd, runc and the kubelet container. To make this process as easy as possible we have created an environment for you to get up and running. There is two differences underlying platforms that are used depending on your OS. If you are using Linux the underlying tech is Docker, If you are using MacOS its vagrant. This is a limitation of the MacOS Kernel and overlayfs.

To get this development environment going run `make shell` below we will explain the subtle differences depending on your underlying OS.

### Linux
Make shell will give you a shell in the $GOPATH from there you can test out your code.

### MacOS
Make shell will land you at `vagrant@kreamlet` so you need to `sudo -i` then `cd /go/src/github.com/puppetlabs/kreamlet`. This is symlinked to your local directory. You can test your code from there. Before you can test your code you will need to start containerd run `./hack/entrypoint.sh` then check that the kubelet container has started by running `ctr -n services.linuxkit c ls`



 
