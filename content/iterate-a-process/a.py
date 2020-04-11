import subprocess
a1 = ['ag', '-V']
o1 = subprocess.Popen(a1, stdout=subprocess.PIPE, text=True)
for s1 in o1.stdout:
   print(s1, end='')
