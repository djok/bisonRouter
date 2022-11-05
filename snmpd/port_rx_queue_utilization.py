import re
from subprocess import Popen, PIPE, STDOUT
import snmp_passpersist as snmp
import subprocess

#
# RCLI field indexes
#
RCLI_ID          = 0
RCLI_CNT         = 5


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
	# lcore
	(2, SNMP_INT_T),
	# port
	(3, SNMP_INT_T),
	# rx_queue
	(4, SNMP_INT_T),
	# empty_polls
	(5, SNMP_INT_T),
	#  pkts/poll
	(6, SNMP_INT_T),
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
	entry_index = 1
	pp.add_int(oid, value)

	rcli_cmd = "sh port rx queue utilization"
	process = subprocess.run(f"sudo /usr/sbin/rcli {rcli_cmd}",
		shell=True, check=True, text=True, capture_output=True)
	for line in process.stdout.split('\n'):
		if not line:
			break
		field_values = line.split("\t")
				
		# skip the header
		if field_values[0] == "lcore":
			continue

		# index
		oid = f"1.{entry_index}"
		snmp_output(pp, SNMP_INT_T, oid, entry_index)

		for i in range(0, RCLI_CNT):
			# skip RCLI fields that doesn't have a corresponding MIB field
			if FIELDS[i][0] == -1:
				continue

			oid = f"{FIELDS[i][0]}.{entry_index}"
			value = field_values[i]
			snmp_type = FIELDS[i][1]

			snmp_output(pp, snmp_type, oid, value)
		
		entry_index += 1


pp = snmp.PassPersist('.1.3.6.1.4.1.2238.240.1.5.1')
pp.start(update, 10)
