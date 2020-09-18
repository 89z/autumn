from functools import reduce
a = ['May', 'June']
# example A
sA = reduce(lambda sc, sd: sc + sd, a)
# example B
f = lambda sc, sd: sc + sd
sB = reduce(f, a)
# print
print(sA == 'MayJune', sB == 'MayJune')
