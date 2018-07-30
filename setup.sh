#!/bin/bash -ex     

make binary

# Chmod the ssh keys
chmod 600 ssh/id_rsa

# Running the instance
if [[ "$OSTYPE" == "linux-gnu" ]]; then
	cd bin/ && ./kream run -c 2 -m 2048 -k 6444
elif [[ "$OSTYPE" == "darwin"* ]]; then
	cd bin/ && ./kream-darwin run -c 2 -m 2048 -k 6444	
fi	
