import os, subprocess
# example 1
os.system('firefox example.com')
# example 2
subprocess.run(['firefox', 'example.com'])
# example 3
subprocess.call(['firefox', 'example.com'])
# example 4
subprocess.check_call(['firefox', 'example.com'])
