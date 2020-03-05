import subprocess
s1 = 'cal'
# example 1
p1 = subprocess.check_output(s1)
s2 = p1.decode()
# example 2
p2 = subprocess.Popen(s1, stdout=subprocess.PIPE)
s3 = p2.stdout.read()
# example 3
p3 = subprocess.Popen(s1, stdout=subprocess.PIPE)
a1 = p3.stdout.readlines()
# print
print(s2, s3, a1)
