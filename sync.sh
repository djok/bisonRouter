#!/bin/bash

rm ./conf/bng-vt-1/brouter.conf
rm ./conf/bng-vt-2/brouter.conf
git pull
../set_brouter.conf.sh
cp ./conf/$HOSTNAME/brouter.conf /etc/bisonrouter/brouter.conf