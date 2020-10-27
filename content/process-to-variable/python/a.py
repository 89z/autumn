import subprocess
a = ['python', '-V']
o = subprocess.Popen(a, stdout=subprocess.PIPE, text=True)
s = o.stdout.read()
print(s, end='')
