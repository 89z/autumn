import subprocess
a = ['python', '-V']
# example 1
o = subprocess.Popen(a, stdout=subprocess.PIPE, text=True)
s1 = o.stdout.read()
# example 2
o = subprocess.run(a, stdout=subprocess.PIPE, text=True)
s2 = o.stdout
# example 3
s3 = subprocess.check_output(a, text=True)
# print
print(s1, s2, s3, end='', sep='')
