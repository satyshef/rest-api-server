#!/bin/bash

##Docker installation script. For ubruntu

if [ -x "$(command -v docker)" ]; then
    echo "Docker installed"
    exit 0
fi


echo "Install docker..."
sudo apt update
sudo apt install -y docker.io
sudo systemctl start docker.service
sudo systemctl enable docker.service
sudo groupadd docker
sudo usermod -aG docker ${USER}



