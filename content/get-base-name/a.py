from os import path
s = '/a/b.tar.xz'
# example 1
s1 = path.basename(s)
# example 2
a = path.splitext(s)
s2 = a[1]
# print
print(s1 == 'b.tar.xz', s2 == '.xz')
