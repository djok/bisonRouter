#!/bin/bash

iptables -t nat -F
iptables -F
iptables -t nat -A PREROUTING -i portal_host -j DNAT -p udp --dport 53 --to-destination {{if .Master}}{{.if.host_mng.ip_master}}{{else if .Backup}}{{.if.host_mng.ip_backup}}{{end}}:53
iptables -t nat -A PREROUTING -i portal_host -j DNAT -p tcp --dport 80 --to-destination {{.tng.portal_server}}:80
iptables -t nat -A PREROUTING -i portal_host -j DNAT -p tcp --dport 443 --to-destination {{.tng.portal_server}}:443
iptables -A FORWARD -i portal_host -j ACCEPT -p tcp --dport 80 -d {{.tng.portal_server}}
iptables -A FORWARD -i portal_host -j ACCEPT -p tcp --dport 443 -d {{.tng.portal_server}}
iptables -A FORWARD -i portal_host -j ACCEPT -p udp --dport 53
iptables -A FORWARD -i portal_host -j ACCEPT -p icmp
iptables -A FORWARD -i portal_host -j DROP
iptables -A INPUT -i portal_host -j DROP
