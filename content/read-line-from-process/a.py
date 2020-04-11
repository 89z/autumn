import subprocess
a1 = ['ag', '-V']
o1 = subprocess.Popen(a1, stdout=subprocess.PIPE, text=True)
s1 = o1.stdout.readline()
print(s1)
