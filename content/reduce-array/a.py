from functools import reduce
a = ['May', 'June']
# example 1
s1 = reduce(lambda s, s1: s + s1, a)
# example 2
f = lambda s, s2: s + s2
s2 = reduce(f, a)
# print
print(s1 == 'MayJune', s2 == 'MayJune')
