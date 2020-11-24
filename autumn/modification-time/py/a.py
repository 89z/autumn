import os, time
# example 1
n1 = time.time()
os.utime('a.py', (n1, n1))
# example 2
n2 = os.path.getmtime('a.py')
# print
print(n1 == n2)
