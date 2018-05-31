#!/bin/bash -ex 
vagrant up
vagrant halt
vagrant package --output ~/Desktop/kreamlet.box
