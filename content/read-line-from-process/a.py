import subprocess
a1 = ['ag', '-V']
o1 = subprocess.Popen(a1, stdout=subprocess.PIPE)
o2 = o1.stdout
a2 = o2.readline()
print(a2)
