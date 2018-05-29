#!/bin/bash -ex     
# Running updates and getting git and wget
sed -i 's/http:\/\//http:\/\/au./g' /etc/apt/sources.list
apt-get update -y
apt-get install -y wget git
# Adding kubernetes.default to the host file for kubectl to us
echo -e "127.0.0.1 kubernetes.default">>/etc/hosts
# Adding ssh keys to connect to Linuxkit
if [ ! -d /root/.ssh ]; then 
	mkdir /root/.ssh
fi       	
cp /vagrant/ssh/id_rsa /root/.ssh && chmod 600 /root/.ssh/id_rsa    
# Installing go if its not there
if [ ! -d /usr/local/go ]; then 
	wget https://dl.google.com/go/go1.9.4.linux-amd64.tar.gz  && tar -C /usr/local -xzf go1.9.4.linux-amd64.tar.gz
fi
# Getting Linxkit 
if [ ! -f /usr/bin/linuxkit ]; then  
	/usr/local/go/bin/go get -u github.com/linuxkit/linuxkit/src/cmd/linuxkit  
	ln -s /root/go/bin/linuxkit /usr/bin/linuxkit
fi       	
# Copying over source code
mkdir -p /root/go/src/github.com/scotty-c/kream-v2
cp -nR /vagrant/* /root/go/src/github.com/scotty-c/kream-v2
# Installing Docker 
if [ ! -f /usr/bin/docker ]; then 
	sudo apt-get install apt-transport-https ca-certificates curl software-properties-common -y
	curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
	add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" 
	apt-get update -y && apt-get install docker-ce -y
fi	
# Getting the kube-master.iso from s3
if [ ! -d ~/.kream/ ]; then 
	mkdir -p ~/.kream
	wget https://s3.amazonaws.com/puppet-cloud-and-containers/kream/kube-master.iso -O ~/.kream/kube-master.iso 
fi
# Making kream binary
cd /root/go/src/github.com/scotty-c/kream-v2 && make binary
# Symlinking kream binary to /usr/bin
if [ ! -f /usr/bin/kream ]; then  
	ln -s /root/go/src/github.com/scotty-c/kream-v2/bin/kream /usr/bin/kream 
fi
# Installing kubectl
if [ ! -f /usr/local/bin/kubectl ]; then 
	curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
	chmod +x ./kubectl && mv ./kubectl /usr/local/bin/kubectl
fi
# Checking to see if the instance is already running, if not create it
if docker ps | grep 6443 ; then
        echo "Linuxkit is running"
else	
	kream run -c 2 -m 2048 -s 3333
fi      	
