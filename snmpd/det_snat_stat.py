import re
from subprocess import Popen, PIPE, STDOUT
import snmp_passpersist as snmp
import subprocess


#
# RCLI field indexes
#
RCLI_VIF_ID      = 0
RCLI_CNT         = 9


#
# SNMP types
#
SNMP_INT_T       = 1
SNMP_CNT64_T     = 2
SNMP_STR_T       = 3
SNMP_IP_T        = 4
SNMP_OCTET_T     = 5


#
# Field metadata: name, OID, type
#
FIELDS = (
	("total maps",            1, SNMP_INT_T),
	("total sessions",        2, SNMP_CNT64_T),
	("portmap failures",      3, SNMP_CNT64_T),
	("portmap failures2",     4, SNMP_CNT64_T),
	("sess overflow",         5, SNMP_CNT64_T),
	("no free portmap ports", 6, SNMP_CNT64_T),
)

FLD_CNT = 6


def find_field(fld_name):
	for i in range(0, FLD_CNT):
		if FIELDS[i][0] == fld_name:
			return i
	return -1


def snmp_output(pp, snmp_type, oid, value):
	if snmp_type == SNMP_INT_T:
		pp.add_int(oid, value)
	elif snmp_type == SNMP_CNT64_T:
		pp.add_cnt_64bit(oid, value)
	elif snmp_type == SNMP_STR_T:
		pp.add_str(oid, value)
	elif snmp_type == SNMP_IP_T:
		pp.add_ip(oid, value)
	elif snmp_type == SNMP_OCTET_T:
		pp.add_oct(oid, parse_mac_addr(value))


def update():
	rcli_cmd = "sh det snat stat"
	process = subprocess.run(f"sudo /usr/sbin/rcli {rcli_cmd}",
		shell=True, check=True, text=True, capture_output=True)
	for line in process.stdout.split('\n'):
		if not line:
			break
		cols = line.split(":")

		i = find_field(cols[0])
		if (i == -1):
			continue

		oid = f"{FIELDS[i][1]}.0"
		value = cols[1].strip()
		snmp_type = FIELDS[i][2]

		snmp_output(pp, snmp_type, oid, value)


pp = snmp.PassPersist('.1.3.6.1.4.1.2238.240.1.4')
pp.start(update, 10)
