import subprocess
a = ['ag', '-V']
o = subprocess.Popen(a, stdout=subprocess.PIPE, text=True)

for s in o.stdout:
   print(s, end='')
