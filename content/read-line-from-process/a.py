import subprocess
p1 = subprocess.Popen('cal', stdout=subprocess.PIPE)
s1 = p1.stdout.readline()
print(s1)
