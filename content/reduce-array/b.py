def f(sa, sc):
   return sa + sc

a = ['May', 'June']
s = ''

for sc in a:
   s = f(s, sc)

print(s == 'MayJune')
