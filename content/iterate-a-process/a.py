import subprocess
a1 = ['ag', '-V']
# example 1
o1 = subprocess.Popen(a1, stdout=subprocess.PIPE, text=True)
for s1 in o1.stdout:
   print(s1, end='')
# example 2
o2 = subprocess.Popen(a1, stdout=subprocess.PIPE)
for a2 in o2.stdout:
   s2 = a2.decode()
   print(s2, end='')
