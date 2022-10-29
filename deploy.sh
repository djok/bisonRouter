#!/bin/bash

# install renderizer if not installed 
if [[ ! -f "./renderizer" ]]; then
  wget https://github.com/gomatic/renderizer/releases/download/v2.0.13/renderizer_linux_x86_64.tar.gz && tar zxvf renderizer_linux_x86_64.tar.gz renderizer && rm renderizer_linux_x86_64.tar.gz
fi

# install keepalived if not installed
if [[ ! -f "/etc/keepalived/keepalived.conf" ]]; then
    apt install keepalived -y
    sed -i 's/#net.ipv4.ip_forward=1/net.ipv4.ip_forward=1/' /etc/sysctl.conf
    if ! grep "net.ipv4.ip_nonlocal_bind=1" /etc/sysctl.conf; then
        echo "net.ipv4.ip_nonlocal_bind=1" >> /etc/sysctl.conf
    fi
    touch /etc/keepalived/keepalived.conf
    sysctl -p
    systemctl enable --now keepalived
fi

# do backup of current config
ts=$(date '+%Y%m%dT%H%M%S')
cp /etc/bisonrouter/brouter.conf /etc/bisonrouter/brouter.conf.$ts
cp /etc/netplan/00-installer-config.yaml /etc/netplan/00-installer-config.yaml.$ts
cp /etc/keepalived/keepalived.conf /etc/keepalived/keepalived.conf.$ts


# select master or backup router
if [[ $(hostname) eq "bng-vt-1" ]]; then ROLE="master"; else ROLE="backup" fi

# generate new config and rollback if anything fails
if ! ./renderizer ./tmpl/brouter.conf --settings=brouter.yaml --master=true --missing zero > /etc/bisonrouter/brouter.conf; then
    echo error in /tmpl/brouter.conf
    echo Rolling back configuration
    mv /etc/bisonrouter/brouter.conf.$ts /etc/bisonrouter/brouter.conf
    mv /etc/netplan/00-installer-config.yaml.$ts /etc/netplan/00-installer-config.yaml
    mv /etc/keepalived/keepalived.conf.$ts /etc/keepalived/keepalived.conf
elif ! ./renderizer ./tmpl/00-installer-config.yaml --settings=brouter.yaml --master=true --missing zero > /etc/netplan/00-installer-config.yaml; then 
    echo error in /tmpl/00-installer-config.yaml
    echo Rolling back configuration
    mv /etc/bisonrouter/brouter.conf.$ts /etc/bisonrouter/brouter.conf
    mv /etc/netplan/00-installer-config.yaml.$ts /etc/netplan/00-installer-config.yaml
    mv /etc/keepalived/keepalived.conf.$ts /etc/keepalived/keepalived.conf
elif ! ./renderizer ./tmpl/keepalived.conf --settings=brouter.yaml --master=true --missing zero > /etc/keepalived/keepalived.conf; then 
    echo error in /tmpl/keepalived.conf
    echo Rolling back configuration
    mv /etc/bisonrouter/brouter.conf.$ts /etc/bisonrouter/brouter.conf
    mv /etc/netplan/00-installer-config.yaml.$ts /etc/netplan/00-installer-config.yaml
    mv /etc/keepalived/keepalived.conf.$ts /etc/keepalived/keepalived.conf
fi

