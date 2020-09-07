import subprocess
a = ['python', '-V']
# example 1
subprocess.call(a)
# example 2
subprocess.check_call(a)
# example 3
subprocess.run(a)
