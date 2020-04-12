import subprocess
s1 = subprocess.check_output('ag', text=True)
print(s1)
