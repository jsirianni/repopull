#!/bin/bash
cd $(dirname $0)

sudo git pull
sudo git status
read -p "Press [Enter] key to upgrade repopull on this branch"
sudo service repopll stop
sudo rm /usr/local/bin/repopull
sudo cp ./bin/repopull /usr/local/bin/
sudo service repopull start
sudo service repopull status
