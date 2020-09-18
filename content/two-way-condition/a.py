n = 10
# example A
if n > 0:
   sA = '+'
elif n < 0:
   sA = '-'
else:
   sA = 'zero'
# example B
sB = '+' if n > 0 else '-'
# print
print(sA == '+', sB == '+')
