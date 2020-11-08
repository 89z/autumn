import os, time
s = 'index.md'
# example 1
n1 = time.time()
os.utime(s, (n1, n1))
# example 2
n2 = os.path.getmtime(s)
# print
print(n1 == n2)
