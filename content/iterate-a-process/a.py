import subprocess
p1 = subprocess.Popen('cal', stdout=subprocess.PIPE, text=True)
for s1 in p1.stdout:
   print(s1, end='')
