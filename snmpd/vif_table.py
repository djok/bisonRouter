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
# Field metadata: OID, type
#
FIELDS = (
	# RCLI_VIF_ID
	(1, SNMP_INT_T),
	# RCLI_VIF_NAME
	(2, SNMP_STR_T),
	# RCLI_PORT
	(3, SNMP_INT_T),
	# RCLI_SVID
	(4, SNMP_INT_T),
	# RCLI_CVID
	(5, SNMP_INT_T),
	# RCLI_RX_PKTS
	(6, SNMP_CNT64_T),
	# RCLI_TX_PKTS
	(7, SNMP_CNT64_T),
	# RCLI_RX_BYTES
	(8, SNMP_CNT64_T),
	# RCLI_TX_BYTES
	(9, SNMP_CNT64_T),
)


def parse_mac_addr(mac):
	res = re.match('^((?:(?:[0-9a-f]{2}):){5}[0-9a-f]{2})$', mac.lower())
	if res is None:
		return '00 00 00 00 00 00'
	return res.group(0).replace(':', ' ')


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
	# this oid is needed to snmptable to work properly
	oid = "0"
	value = 0
	pp.add_int(oid, value)

	rcli_cmd = "sh vif counters"
	process = subprocess.run(f"sudo /usr/sbin/rcli {rcli_cmd}",
		shell=True, check=True, text=True, capture_output=True)
	for line in process.stdout.split('\n'):
		if not line:
			break
		field_values = line.split("\t")

		# skip the header
		if field_values[RCLI_VIF_ID] == "vif_id":
			continue

		for i in range(0, RCLI_CNT):
			# skip RCLI FIELDS that doesn't have a corresponding MIB field
			if FIELDS[i][0] == -1:
				continue

			oid = f"{FIELDS[i][0]}.{int(field_values[RCLI_VIF_ID])}"
			value = field_values[i]
			snmp_type = FIELDS[i][1]

			snmp_output(pp, snmp_type, oid, value)


pp = snmp.PassPersist('.1.3.6.1.4.1.2238.240.1.3.1')
pp.start(update, 10)
