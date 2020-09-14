n = 10
# example 1
if n > 0:
   s = '+'
elif n < 0:
   s = '-'
else:
   s = 'zero'
# example 2
s2 = '+' if n > 0 else '-'
# print
print(s == '+', s2 == '+')
