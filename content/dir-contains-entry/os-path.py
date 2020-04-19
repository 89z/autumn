from os import path
s1 = 'a.txt'
# example 1
b1 = path.exists(s1)
# example 2
b2 = path.isfile(s1)
# print
print(b1, b2)
