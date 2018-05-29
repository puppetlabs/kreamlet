#!/bin/bash -ex     
# Adding kubernetes.default to the host file for kubectl to us
if cat /etc/hosts | grep kubernetes.default  
then
	echo "hostfile entry is already there"
else	
	sudo echo -e "127.0.0.1 kubernetes.default">>/etc/hosts
fi	
# Getting Linxkit 
if [ ! -f $GOPATH/bin/linuxkit ]; then  
	go get -u github.com/linuxkit/linuxkit/src/cmd/linuxkit  
fi       	
# Getting the kube-master.iso from s3
if [ ! -f ~/.kream/kube-master.iso ]; then 
	mkdir -p ~/.kream
	wget https://s3.amazonaws.com/puppet-cloud-and-containers/kream/kube-master.iso -O ~/.kream/kube-master.iso 
fi
# Making kream binary
make binary

# Chmod the ssh keys
chmod 600 ssh/id_rsa

# Running the instance
if [[ "$OSTYPE" == "linux-gnu" ]]; then
	cd bin/ && ./kream run -c 2 -m 2048 -k 6444
elif [[ "$OSTYPE" == "darwin"* ]]; then
	cd bin/ && ./kream-darwin run -c 2 -m 2048 -k 6444	
fi	
