from os import path
s1 = '/sunday/monday.tar.xz'
# example 1
s2 = path.basename(s1)
# example 2
a1 = path.splitext(s1)
s3 = a1[1]
# print
print(s2 == 'monday.tar.xz', s3 == '.xz')
