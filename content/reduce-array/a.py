import functools
a = ['May', 'June']
f = lambda s_acc, s_cur: s_acc + s_cur
s = functools.reduce(f, a)
print(s == 'MayJune')
