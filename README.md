# kream-v2 - Note this project is still in alpha

### Install dependencies
```
Linux:
docker-ce
golang 1.10.2

MacOS:
docker-ce
golang 1.10.2

```

## Getting it up and running 


Firstly you need to get the code with `go get -u github.com/puppetlabs/kreamlet`

Then just run `./setup.sh`

This will boot the Linuxkit instance and make sure you have all the parts you need. 
 

## What is this repo actually doing ?
This is the proof of concept to prove we can use Linuxkit as the mechanism to build and provision a kubernetes OS. To test this theroy we are running linuxkit inside a qemu container via an iso on linux and on mac OS we will use hyperkit.
From the same build we can create any cloud image. From the start of this project we will use cri-containerd instead if Docker. This will greatly reduce our image size.

## Limitations 
At the moment you will have to manually expose services through `kubectl port-forward` This is going to be added to the beta release. 

It only supports Kubernetes v1.10.3.


## Road Map (For this repo)
Here is the functionality that is on the road map for this repo (in order of importance)
 - [ ] Add worker nodes
 - [ ] Add a download mechanism to get the latest iso we create
 - [ ] Add the build jobs/code to this repo
 - [ ] Add ssl to all grpc connections
