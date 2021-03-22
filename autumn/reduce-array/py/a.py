from functools import reduce
a = ['May', 'June']

# example 1
s = reduce(lambda s, t: s + t, a)
print(s == 'MayJune')

# example 2
n = reduce(lambda n, s: n + len(s), a, 0)
print(n == 7)
