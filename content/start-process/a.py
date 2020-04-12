import subprocess
a1 = ['ag', '-V']
# example 1
subprocess.call(a1)
# example 2
subprocess.check_call(a1)
# example 3
subprocess.run(a1)
