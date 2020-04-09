import subprocess
# example 1
r1 = subprocess.Popen('locale', stdout=subprocess.PIPE, text=True)
for s1 in r1.stdout:
   print(s1, end='')
# example 2
r1 = subprocess.Popen('locale', stdout=subprocess.PIPE)
for a1 in r1.stdout:
   s1 = a1.decode()
   print(s1, end='')
