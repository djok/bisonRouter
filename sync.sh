#!/bin/bash

rm  ./brouter.conf
git pull
./set_brouter.conf.sh
cp ./brouter.conf /etc/bisonrouter/brouter.conf