from functools import reduce
a = ['May', 'June']
# example 1
s = reduce(lambda s, s1: s + s1, a)
# example 2
n = reduce(lambda n, s: n + len(s), a, 0)
# print
print(s == 'MayJune', n == 7)
