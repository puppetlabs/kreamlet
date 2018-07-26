#!/bin/bash -ex     
# Adding kubernetes.default to the host file for kubectl to us
if cat /etc/hosts | grep kubernetes.default  
then
	echo "hostfile entry is already there"
else	
	sudo echo "127.0.0.1 kubernetes.default">>/etc/hosts
fi	
# Getting Linxkit 
if [ ! -f $GOPATH/bin/linuxkit ]; then  
	go get -u github.com/linuxkit/linuxkit/src/cmd/linuxkit  
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
