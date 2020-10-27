import subprocess
a = ['python', '-V']
o = subprocess.run(a, stdout=subprocess.PIPE, text=True)
print(o.stdout, end='')
