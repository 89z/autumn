import os
s = 'USERPROFILE'
# example 1
s2 = os.getenv(s)
# example 2
s3 = os.environ[s]
# print
print(s2, s3, sep='\n')
