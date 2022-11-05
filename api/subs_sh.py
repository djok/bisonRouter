from ttp import ttp
import subprocess

ttp_template = """{{ vif_id }} {{ port | DIGIT }} {{ vlan }} {{ ip | IP }} {{ mac | MAC }} {{ sess-id }} {{ circuit-id }} {{ remote-id }} {{ ingress_car }} {{ egress_car }} {{ rx_pkts }} {{ tx_pkts }} {{ rx_bytes }} {{ tx_bytes }} {{ pbr }} {{ l2_pbr }} {{ ttl }} {{ expire_in }} {{ uptime }}"""

process = subprocess.run(['rcli', 'sh subs'], 
                         stdout=subprocess.PIPE, 
                         universal_newlines=True)
# process

# create parser object and parse data using template:
parser = ttp(data=process.stdout, template=ttp_template)
parser.parse()

# print result in JSON format
results = parser.result(format='json')[0]
print(results)