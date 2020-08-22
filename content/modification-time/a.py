import os, time
n1 = time.time()
os.utime('index.md', (n1, n1))
n2 = os.path.getmtime('index.md')
print(n1, n2 == n1)
