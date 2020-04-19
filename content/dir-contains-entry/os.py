import os
s1 = 'a.txt'
b1 = os.access(s1, os.F_OK)
print(b1)
