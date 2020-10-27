import subprocess
a = ['python', '-V']
s = subprocess.check_output(a, text=True)
print(s, end='')
