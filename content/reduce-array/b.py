a = ['May', 'June']
f = lambda s_acc, s_cur: s_acc + s_cur
s = ''

for s_cur in a:
   s = f(s, s_cur)

print(s == 'MayJune')
