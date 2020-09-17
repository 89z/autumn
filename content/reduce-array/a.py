from functools import reduce
a = ['May', 'June']
# example 1
s = reduce(lambda sa, sc: sa + sc, a)
# example 2
f = lambda sa, sc: sa + sc
s2 = reduce(f, a)
# print
print(s == 'MayJune', s2 == 'MayJune')
