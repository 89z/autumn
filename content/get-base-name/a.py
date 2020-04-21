from os import path
s1 = '/sunday/monday.tar.xz'
# example 1
s2 = path.basename(s1)
b1 = s2 == 'monday.tar.xz'
# example 2
a1 = path.splitext(s1)
s1 = a1[1]
b2 = s1 == '.xz'
# print
print(b1, b2)
