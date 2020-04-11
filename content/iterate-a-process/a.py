import subprocess
a1 = ['ag', '-V']
# example 1
o1 = subprocess.Popen(a1, stdout=subprocess.PIPE, text=True)
for s1 in o1.stdout:
   print(s1, end='')
# example 2
o2 = subprocess.Popen(a1, stdout=subprocess.PIPE)
for a1 in o2.stdout:
   print(a1)
