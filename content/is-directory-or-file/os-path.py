from os import path
s = 'index.md'
# example 1
b = path.isfile(s)
# example 2
b2 = path.exists(s)
# example 3
b3 = not path.isdir(s)
# print
print(b, b2, b3)
