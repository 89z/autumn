import subprocess
a1 = ['ag', '-V']
# example 1
s1 = subprocess.check_output(a1, text=True)
# example 2
o1 = subprocess.Popen(a1, stdout=subprocess.PIPE, text=True)
s2 = o1.stdout.read()
# example 3
o3 = subprocess.Popen(a1, stdout=subprocess.PIPE, text=True)
a2 = o3.stdout.readlines()
# print
print(s1, s2, a2, sep='\n')
