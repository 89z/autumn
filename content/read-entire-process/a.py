import subprocess
a1 = ['ag', '-V']
# example 1
a2 = subprocess.check_output(a1)
# example 2
o1 = subprocess.Popen(a1, stdout=subprocess.PIPE)
a3 = o1.stdout.read()
# example 3
o2 = subprocess.Popen(a1, stdout=subprocess.PIPE)
a4 = o2.stdout.readlines()
# print
print(a2, a3, a4, sep='\n')
