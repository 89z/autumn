import os
s = 'USERPROFILE'
# example 1
s1 = os.getenv(s)
# example 2
s2 = os.environ[s]
# print
s = r'C:\Users\Steven'
print(s1 == s, s2 == s)
