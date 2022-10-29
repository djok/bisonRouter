# bisonRouter


VRRP Setup
add "net.ipv4.ip_nonlocal_bind=1" 
sudo nano /etc/sysctl.conf
```
net.ipv4.ip_nonlocal_bind=1
```
run sudo sysctl -p
sudo apt-get install keepalived
sudo nano /etc/keepalived/keepalived.conf

sudo service keepalived start
sudo systemctl status keepalived
sudo service keepalived stop

# install template parser

mkdir renderizer && cd renderizer && wget https://github.com/gomatic/renderizer/releases/download/v2.0.13/renderizer_linux_x86_64.tar.gz && tar zxvf renderizer_linux_x86_64.tar.gz && sudo mv ./renderizer /usr/bin/. && cd ..

```
#!/bin/bash
ts=$(date '+%Y%m%dT%H%M%S')
cp /etc/bisonrouter/brouter.conf /etc/bisonrouter/brouter.conf.$ts
cp /etc/netplan/00-installer-config.yaml /etc/netplan/00-installer-config.yaml.$ts
cp /etc/keepalived/keepalived.conf /etc/keepalived/keepalived.conf.$ts


if ! renderizer ./tmpl/brouter.conf --settings=brouter.yaml --master=true --missing zero > /etc/bisonrouter/brouter.conf
then
    echo error in /tmpl/brouter.conf
    echo Rolling back configuration
    cp /etc/bisonrouter/brouter.conf.$ts /etc/bisonrouter/brouter.conf
    cp /etc/netplan/00-installer-config.yaml.$ts /etc/netplan/00-installer-config.yaml
    cp /etc/keepalived/keepalived.conf.$ts /etc/keepalived/keepalived.conf
elseif ! renderizer ./tmpl/00-installer-config.yaml --settings=brouter.yaml --master=true --missing zero > /etc/netplan/00-installer-config.yaml
then 
    echo error in /tmpl/00-installer-config.yaml
    echo Rolling back configuration
    cp /etc/bisonrouter/brouter.conf.$ts /etc/bisonrouter/brouter.conf
    cp /etc/netplan/00-installer-config.yaml.$ts /etc/netplan/00-installer-config.yaml
    cp /etc/keepalived/keepalived.conf.$ts /etc/keepalived/keepalived.conf
elseif ! ./tmpl/keepalived.conf --settings=brouter.yaml --master=true --missing zero > /etc/keepalived/keepalived.conf
then 
    echo error in /tmpl/keepalived.conf
    echo Rolling back configuration
    cp /etc/bisonrouter/brouter.conf.$ts /etc/bisonrouter/brouter.conf
    cp /etc/netplan/00-installer-config.yaml.$ts /etc/netplan/00-installer-config.yaml
    cp /etc/keepalived/keepalived.conf.$ts /etc/keepalived/keepalived.conf
fi
```

