n = 10

# example 1
if n > 0:
   s1 = '+'
elif n < 0:
   s1 = '-'
else:
   s1 = 'zero'

# example 2
s2 = '+' if n > 0 else '-'

# print
print(s1 == '+', s2 == '+')
