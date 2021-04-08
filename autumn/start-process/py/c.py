import subprocess

# example 1
subprocess.run('dust -V')

# example 2
subprocess.run(['dust', '-V'])

# example 3
p = subprocess.run('dust -V', capture_output=True)
print(p.stdout)
