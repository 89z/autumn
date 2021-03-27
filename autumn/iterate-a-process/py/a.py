import subprocess
a = ['python', '-h']

p = subprocess.Popen(a, stdout=subprocess.PIPE, text=True)
for s in p.stdout:
   print(s, end='')
