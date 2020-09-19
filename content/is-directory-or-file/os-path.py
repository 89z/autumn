from os import path
s = 'index.md'
# example 1
b1 = path.isfile(s)
# example 2
b2 = path.exists(s)
# example 3
b3 = not path.isdir(s)
# print
print(b1, b2, b3)
