from ttp import ttp
import subprocess

ttp_template = """{{ port | DIGIT }} {{ vid }} {{ ip | IP }} {{ mac | MAC }} {{ type }} {{ state }}"""


process = subprocess.run(['rcli', 'sh arp cache'], 
                         stdout=subprocess.PIPE, 
                         universal_newlines=True)
# process

# create parser object and parse data using template:
parser = ttp(data=process.stdout, template=ttp_template)
parser.parse()

# print result in JSON format
results = parser.result(format='json')[0]
print(results)