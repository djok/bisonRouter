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

if [[ ! -f "/etc/bisonrouter/iptables.sh" ]]; then
    touch /etc/bisonrouter/iptables.sh
fi


if [[ ! -f "/etc/dnsmasq.conf" ]]; then
    apt install dnsmasq -y
    systemctl enable --now dnsmasq
fi

# do backup of current config
ts=$(date '+%Y%m%dT%H%M%S')
cp /etc/bisonrouter/brouter.conf /etc/bisonrouter/brouter.conf.$ts
cp /etc/bisonrouter/iptables.sh /etc/bisonrouter/iptables.sh.$ts
cp /etc/netplan/00-installer-config.yaml /etc/netplan/00-installer-config.yaml.$ts
cp /etc/keepalived/keepalived.conf /etc/keepalived/keepalived.conf.$ts
cp /etc/dnsmasq.conf /etc/dnsmasq.conf.$ts

roll_back () {
    echo Rolling back configuration
    mv /etc/bisonrouter/brouter.conf.$1 /etc/bisonrouter/brouter.conf
    mv /etc/bisonrouter/iptables.sh.$1 /etc/bisonrouter/iptables.sh
    mv /etc/netplan/00-installer-config.yaml.$1 /etc/netplan/00-installer-config.yaml
    mv /etc/keepalived/keepalived.conf.$1 /etc/keepalived/keepalived.conf
    mv /etc/dnsmasq.conf.$1 /etc/dnsmasq.conf
}

# select master or backup router
if [[ $(hostname) -eq "bng-vt-1" ]]; then ROLE=master; else ROLE=backup; fi

# generate new config and rollback if anything fails
if ! ./renderizer ./tmpl/brouter.conf --settings=brouter.yaml --$ROLE=true --missing zero > /etc/bisonrouter/brouter.conf; then
    echo error in /tmpl/brouter.conf
    roll_back $ts
elif ! ./renderizer ./tmpl/00-installer-config.yaml --settings=brouter.yaml --$ROLE=true --missing zero > /etc/netplan/00-installer-config.yaml; then 
    echo error in /tmpl/00-installer-config.yaml
    roll_back $ts
elif ! ./renderizer ./tmpl/keepalived.conf --settings=brouter.yaml --$ROLE=true --missing zero > /etc/keepalived/keepalived.conf; then 
    echo error in /tmpl/keepalived.conf
    roll_back $ts
elif ! ./renderizer ./tmpl/deploy_iptables.sh --settings=brouter.yaml --$ROLE=true --missing zero > /etc/bisonrouter/iptables.sh; then
    echo error in /tmpl/deploy_iptables.sh
    roll_back $ts
elif ! ./renderizer ./tmpl/dnsmasq.conf --settings=brouter.yaml --$ROLE=true --missing zero > /etc/dnsmasq.conf; then
    echo error in /tmpl/dnsmasq.conf
    roll_back $ts
else
    if ! cmp -s /etc/bisonrouter/brouter.conf.$ts /etc/bisonrouter/brouter.conf; then
        echo new /etc/bisonrouter/brouter.conf compared to /etc/bisonrouter/brouter.conf.$ts
        read -p "Confirm applying brouter.conf config? " -n 1 -r
        echo    # (optional) move to a new line
        if [[ $REPLY =~ ^[Yy]$ ]]
        then
            bisonrouter restart
        fi
    fi
    if ! cmp -s /etc/bisonrouter/iptables.sh.$ts /etc/bisonrouter/iptables.sh; then
        echo new /etc/bisonrouter/iptables.sh compared to /etc/bisonrouter/iptables.sh.$ts
        read -p "Confirm applying iptables.sh? " -n 1 -r
        echo    # (optional) move to a new line
        if [[ $REPLY =~ ^[Yy]$ ]]
        then
            bash /etc/bisonrouter/iptables.sh
        fi
    fi
    if ! cmp -s /etc/netplan/00-installer-config.yaml.$ts /etc/netplan/00-installer-config.yaml; then
        echo new /etc/netplan/00-installer-config.yaml compared to /etc/netplan/00-installer-config.yaml.$ts
        read -p "Confirm applying new netplan config? " -n 1 -r
        echo    # (optional) move to a new line
        if [[ $REPLY =~ ^[Yy]$ ]]
        then
            netplan apply
        fi
    fi
    if ! cmp -s /etc/keepalived/keepalived.conf.$ts /etc/keepalived/keepalived.conf; then
        echo new /etc/keepalived/keepalived.conf compared to /etc/keepalived/keepalived.conf.$ts
        read -p "Confirm applying new keepalived config? " -n 1 -r
        echo    # (optional) move to a new line
        if [[ $REPLY =~ ^[Yy]$ ]]
        then
            systemctl restart keepalived
        fi
    fi
    if ! cmp -s /etc/dnsmasq.conf.$ts /etc/dnsmasq.conf; then
        echo new /etc/dnsmasq.conf compared to /etc/dnsmasq.conf.$ts
        read -p "Confirm applying new dnsmasq config? " -n 1 -r
        echo    # (optional) move to a new line
        if [[ $REPLY =~ ^[Yy]$ ]]
        then
            systemctl restart dnsmasq
        fi
    fi
fi

