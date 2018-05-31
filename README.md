# kream-v2

### Install dependencies
```
Linux:
docker-ce
golang 1.10.2

MacOS:
vagrant 
golang 1.10.2

```

## Getting it up and running 


Firstly you need to get the code with `go get github.com/puppetlabs/kreamlet`

Then just run `./setup.sh`

This will boot the Linuxkit instance and make sure you have all the parts you need. 

After the machine boots use the `ssh.sh` script in the root directory to ssh into the instance.

Once you have the ssh terminal up run `ctr --namespace services.linuxkit tasks exec --tty --exec-id kube kubelet kubeadm-init.sh`

This will boot strap Kubernetes, once the cluster is built run `ctr --namespace services.linuxkit tasks exec --tty --exec-id kube kubelet cat /etc/kubernetes/admin.conf`

Copy the admin.conf to your local machine somewhere then edit the line that has `server: https://<DHCP ADDRESS>:6443` to `server: https://kubernetes.default:6444` run `export KUBECONFIG=admin.conf`

kubectl should be able to hit the instance. 

## What is this repo actually doing ?
This is the proof of concept to prove we can use Linuxkit as the mechanism to build and provision a kubernetes OS. To test this theroy we are running linuxkit inside a qemu container via an iso.
From the same build we can create any cloud image. From the start of this project we will use cri-containerd instead if Docker. This will greatly reduce our image size.

## Limitations 
At the moment this only works on linux at the moment due to Docker not being able to naively on mac and integration we need to build into vpnkit to get networking. We will circle back round to this and use hyperkit once we have all the other functionality in.


## Road Map (For this repo)
Here is the functionality that is on the road map for this repo (in order of importance)
 - [ ] Add build functionality to orchestrate the cluster creation. Initial work for this has started an is located under the bootstrap folder
 - [ ] Add worker nodes
 - [ ] Add Vagrantfile for mac users until hyperkit is available 
 - [ ] Add a download mechanism to get the latest iso we create
 - [ ] Add the build jobs/code to this repo


## Other thoughts
This is going to be donated to Puppet once it is in alpha.
As it will be a mono repo for the k8 platform we should rename it (this can be done on migration to puppet)
