# example 1
a = ['May', 'June']
s = ','.join(a)

# example 2
a = [10, 11]
s2 = ''
for n, n2 in enumerate(a):
   if n > 0: s2 += ','
   s2 += str(n2)

# print
print(s, s2)
