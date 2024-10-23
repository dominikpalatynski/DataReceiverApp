#!/bin/bash

# Update the package list
sudo apt-get update

# Install prerequisites for Docker
sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common

# Add Docker's official GPG key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# Add Docker's official repository to APT sources
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Update the package list again after adding Docker's repository
sudo apt-get update

# Install Docker CE (Community Edition)
sudo apt-get install -y docker-ce

# Add the current user to the docker group
sudo usermod -aG docker ${USER}

# Enable and start the Docker service
sudo systemctl enable docker
sudo systemctl start docker

# Verify Docker installation by printing the version
docker --version
