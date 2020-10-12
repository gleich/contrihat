#!/bin/sh
# Installing docker onto a RPI

# Install Docker:
curl -sSL https://get.docker.com | sh

# Giving Docker Sudo Privildages:
sudo groupadd docker
sudo gpasswd -a $USER docker

# Installing docker compose

sudo apt install docker-compose -y

# Rebooting the pi
echo "The pi will reboot in 5 seconds"
sleep 5
sudo reboot
