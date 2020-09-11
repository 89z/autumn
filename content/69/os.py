import os
s = 'index.md'
b = os.access(s, os.F_OK)
print(b)
