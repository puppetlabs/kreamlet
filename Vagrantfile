# -*- mode: ruby -*-
# # vi: set ft=ruby :
 
# Specify minimum Vagrant version and Vagrant API version
Vagrant.require_version ">= 1.6.0"
VAGRANTFILE_API_VERSION = "2"
 
# Require YAML module
require 'yaml'
 
# Read YAML file with box details
servers = YAML.load_file('bootstrap/servers.yaml')
 
# Create boxes
Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
 
  # Iterate through entries in YAML file
servers.each do |servers|
    
  config.vm.define servers["name"] do |srv|
    
    srv.vm.hostname = servers["name"]
    
    srv.vm.box = servers["box"]
    
    srv.vm.network "private_network", ip: servers["ip"]

  
   servers["forward_ports"].each do |port| 
     srv.vm.network :forwarded_port, guest: port["guest"], host: port["host"]
  end

   srv.vm.provider :virtualbox do |v|
        v.cpus = servers["cpu"]
        v.memory = servers["ram"]
  end
   
    srv.vm.synced_folder "./", "/home/vagrant/#{servers['name']}"
    
    servers["shell_commands"].each do |sh|
      srv.vm.provision "shell", inline: sh["shell"]
    end
    
      end
    end
  end
