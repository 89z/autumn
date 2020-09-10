import os, time
n = time.time()
os.utime('index.md', (n, n))
n2 = os.path.getmtime('index.md')
print(n, n2 == n)
