# example 1
a = ['May', 'June']
s = ','.join(a)

# example 2
a = [10, 11]
s2 = ''
for n in a:
   if s2 > '': s2 += ','
   s2 += str(n)

# print
print(s == 'May,June', s2 == '10,11')
