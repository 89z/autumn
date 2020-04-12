import subprocess
a1 = ['ag', '-V']
# example 1
o1 = subprocess.Popen(a1, stdout=subprocess.PIPE, text=True)
o2 = o1.stdout
s1 = o2.read()
# example 2
o3 = subprocess.run(a1, stdout=subprocess.PIPE, text=True)
s2 = o3.stdout
# example 3
s3 = subprocess.check_output(a1, text=True)
# print
print(s1, s2, s3, end='', sep='')
