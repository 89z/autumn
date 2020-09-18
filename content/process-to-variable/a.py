import subprocess
a = ['python', '-V']
# example A
o = subprocess.Popen(a, stdout=subprocess.PIPE, text=True)
sA = o.stdout.read()
# example B
o = subprocess.run(a, stdout=subprocess.PIPE, text=True)
sB = o.stdout
# example C
sC = subprocess.check_output(a, text=True)
# print
print(sA, sB, sC, end='', sep='')
