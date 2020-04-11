import subprocess
a1 = ['ag', '-V']
# example 1
a2 = subprocess.check_output(a1)
s1 = a2.decode()
# example 2
o1 = subprocess.Popen(a1, stdout=subprocess.PIPE)
o2 = o1.stdout
s2 = o2.read().decode()
# example 3
o3 = subprocess.Popen(a1, stdout=subprocess.PIPE)
a2 = o3.stdout.readlines()
# print
print(s1, s2, a2)
